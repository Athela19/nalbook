# TRD Backend — Nalbook

> Implementasi teknis dari [PRD Backend](./prd.md). Skema data: [ERD](./erd.md). API: `api/`.

## 1. Stack & Versi Baseline

| Lapisan | Pilihan | Catatan |
|---|---|---|
| Bahasa | Go 1.22+ (module, `gofmt`/`go vet` di CI) | 1 orang: stdlib + sedikit lib matang |
| HTTP | `chi` + middleware (`auth, tenant, rbac, ratelimit, idempotency`) | OpenAPI 3.1 sebagai kontrak (`api/*.md` → `openapi.yaml`) |
| DB | PostgreSQL 16, driver `pgx/v5` | Satu DB per tenant, `NUMERIC(19,4)` untuk uang |
| Migrasi | `golang-migrate` / `goose`, runner multi-DB + `schema_version` per tenant | Backward-compatible |
| Worker | Go (`asynq` atau `river` di atas Postgres agar tanpa Redis) | Laporan, import, susut, kurs, webhook, backup |
| File | S3-compatible (MinIO lokal / AWS / IDCloud) terenkripsi SSE | Lampiran + hasil export |
| Auth | JWT access (15 mnt) + refresh rotating (httpOnly), `argon2id`, TOTP (`pquerna/otp`), OIDC (`zitadel/oidc`) | SSO OIDC/SAML via broker |
| Config | `env` + vault (tanpa secret di kode), validasi saat boot | `DATABASE_ADMIN_URL, TENANT_DSN_TPL, S3_*, KMS_*, OIDC_*` |
| Infra | Docker Compose (dev) → Docker image + CI/CD (GitHub Actions: lint, test, scan, migrate dry-run) | Server disarankan Indonesia |
| Observabilitas | `slog` JSON + OpenTelemetry trace `trace_id` per request | Tanpa PII di log |

Struktur repo yang disarankan:

```
cmd/api/main.go cmd/worker/main.go cmd/migrate/main.go
internal/{auth,tenant,coa,journal,sales,purchase,cashbank,inventory,asset,tax,budget,project,production,payroll,currency,period,report,import,webhook,audit}/
pkg/{money,authz,pagination,idempotency,openapi}/
migrations/NNNN_name.sql
```

## 2. Konvensi API (mengikat PRD §8)

- Base ` /api/v1`, versioning URL. Format error baku bahasa Indonesia:
  ```json
  { "code": "VALIDATION_ERROR", "message": "Total debit harus sama dengan kredit.", "details": {...}, "trace_id": "..." }
  ```
- Header konteks: `X-Company-ID`, `X-Branch-ID`. Semua query wajib `WHERE tenant_id=? AND company_id=?`.
- Pagination cursor: `?cursor=&limit=` → `{ data, next_cursor }`. Daftar besar selalu server-side.
- `Idempotency-Key` wajib untuk `POST` transaksi (invoice, payment, journal). Duplikat → kembalikan hasil asli.
- OpenAPI digenerate dari handler (mis. `huma`/`oas`) agar frontend & publik selalu setara (PRD §8: API internal = publik).
- Auth: `Authorization: Bearer` (API key / OAuth2 client-credentials untuk integrasi) + cookie session untuk browser (`SameSite=Lax`, `Secure`, CSRF token ganda).

Rate-limit per tenant+key (mis. 600 req/mnt default, endpoint auth lebih ketat + lockout progresif).

## 3. Multi-Tenancy & Migrasi (PRD §3, §9)

- Provisioning `POST /internal/tenants`: buat DB `nalbook_t_<id>`, user DB terpisah, jalankan migrasi, seed template COA per jenis usaha + peran default + `period_locks` awal.
- Koneksi: pool per tenant dengan cache + evict; larangan keras cross-tenant query (test otomatis: cari literal tanpa `tenant_id`).
- Migrasi: `cmd/migrate --all-tenants`, catat `schema_migrations(version, applied_at)` per DB; rollback = migrasi maju perbaikan, bukan down destruktif.
- Backup: `pg_dump` harian per tenant → S3 off-site terenkripsi; drill restore berkala; RPO/RTO mengikuti PRD §11 (open).

## 4. Invarian Akuntansi (PRD §4–§5, diuji otomatis)

1. Transaksi DB atomik per posting; `sum(debit)==sum(credit)` dicek di Go + `CHECK` parsial di DB; tolak jika timpang.
2. Uang: `NUMERIC(19,4)` / minor integer di Go (`shopspring/decimal`), tidak pernah `float64`.
3. Nomor dokumen: tabel `document_sequences(company_id, doc_type, last_no)` dengan `SELECT ... FOR UPDATE` → format `INV/2026/XI/0001`; unique `(company_id, doc_type, doc_no)`.
4. Status `draft → posted → voided`; `posted` immutable (UPDATE ditolak kecuali void/corrective). Void butuh alasan.
5. Kunci periode: tabel `period_locks(company_id, period_start, period_end, is_closed)`; trigger app-level menolak tulis ke periode closed tanpa `period:reopen`.
6. Dimensi: `journal_lines` boleh bawa `project_id, department_id, cost_center_id`, `currency, exchange_rate, rate_date, rate_source`, `tax_code_id`.
7. FIFO: tabel `stock_layers`; rata-rata: `items.avg_cost`; serial/batch: `stock_serials`. Opname → adjustment butuh approve.
8. Susut & kurs & jurnal berulang dijalankan worker idempoten (dedupe key `period+asset_id`).

Audit: setiap mutasi tulis `audit_logs` (append-only, REVOKE UPDATE/DELETE dari role app). Kolom sensitif (NPWP, rekening, gaji) dienkripsi kolom (AES-GCM via KMS/envelope) — lihat ERD.

## 5. RBAC & Keamanan (PRD §9)

- Middleware urutan: `request-id → logger → cors → security-headers → session/apikey → tenant/company scope → rbac(module:action) → idempotency → handler`.
- Matriks `roles → permissions(module, action)` per company, dapat dikustom; konsultan = grant sementara `expires_at`, deny-by-default.
- Password `argon2id`, TOTP 2FA, OIDC/SAML, sesi DB dengan revoke (`sessions` + denylist refresh).
- Validasi input: `go-playground/validator` + decimal bounds; output-encode; `Content-Security-Policy, HSTS, X-Content-Type-Options`; CSRF untuk cookie-auth.
- Secret via env/vault; `govulncheck` + ` Dependabot` di CI; image non-root, read-only FS.
- Test wajib: golden jurnal per modul (banding harapan Accurate sample), concurrent numbering (0 duplikat), cross-tenant leak test, kunci periode test, pentest sebelum rilis.

## 6. Worker & Job (PRD §6, §8)

Antrean (`river`/`asynq`): `report.export, import.commit, depreciation.run, fx.sync, webhook.deliver, backup.run, recurring-journal.run`.

- Export laporan: buat file PDF/Excel → S3 → URL presigned kadaluarsa; `POST /api/v1/reports/{key}/exports`.
- Import Excel/CSV: upload → parse → preview + validasi per baris → `import_batches` → commit atomik; error report per baris.
- Webhook: HMAC-SHA256 signature, retry eksponensial (10x, backoff), `webhook_deliveries` log; event `invoice.posted, payment.received, stock.moved, document.voided`.
- FX sync harian: sumber utama + cadangan (open Q PRD), simpan `exchange_rates`, jangan pernah dummy.

## 7. Performa & Operasional

- Indeks wajib: `(tenant_scoped) (company_id, date)`, `(doc_type, doc_no)`, `(contact_id, due_date)` untuk aging, GIN untuk `audit_logs(new_values)`.
- Target awal (final di PRD open Q): p95 baca <300 ms, posting <800 ms, laporan berat via job async; pagination default 25–50.
- Health `/healthz`, readiness cek DB+S3; metrik Prometheus (request, job, backup); alert aktivitas curiga + gagal login beruntun.

## 8. Keterlacakan PRD → TRD

| PRD § | TRD § |
|---|---|
| §4 kontrak modul | §4 invarian |
| §5.1 kurs, §5.2 tutup buku | §4.7–4.8, §6 FX |
| §6 laporan | §6 export job |
| §7 audit/UU PDP | §4 audit + §5 enkripsi |
| §8 API/webhook/migrasi | §2 + §6 |
| §9 NFR | §5 + §7 |
| §10 arsitektur | §1 + struktur repo |
