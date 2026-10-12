# PRD Backend — Nalbook

> Dokumen ini adalah sisi backend dari PRD Nalbook. Untuk sisi antarmuka lihat [PRD Frontend](../frontend/prd.md).
> Sumber: PRD Nalbook 12 Okt 2026 oleh Muhammad Syarif Nurrohman. Bahasa: Indonesia saja.

## 1. Konteks & Ruang Lingkup Backend

Nalbook adalah SaaS akuntansi web untuk UMKM dan kantor akuntan Indonesia, setara fungsi Accurate dengan integrasi mudah. Target rilis production-ready penuh 12 bulan (12 Okt 2026 → 12 Okt 2027), tim 1 orang, >1.000 pengguna.

| Item | Keputusan Backend |
|---|---|
| Backend | Go, REST API + OpenAPI, dipakai frontend & publik (sama) |
| Frontend | Next.js (dok terpisah), hanya bicara ke backend |
| Isolasi tenant | Database PostgreSQL terpisah per tenant |
| Uang | Desimal presisi tetap / integer minor, tidak pernah float |
| Jurnal | Satu-satunya sumber kebenaran angka |
| Infra | Docker, CI/CD, worker Go, object storage S3-compatible, server disarankan Indonesia |
| Bahasa | Indonesia saja, format tanggal/angka Indonesia di serialisasi |

In-scope backend: seluruh §5–§8 PRD induk (17 modul P0, laporan sebagai query jurnal, pajak Indonesia, API/webhook/integrasi, import Excel/CSV, export PDF/Excel, 2FA/SSO, enkripsi, backup, audit log).

Out-of-scope rilis ini: billing/langganan di dalam sistem, aplikasi native, mode offline, bahasa selain Indonesia, import format software lain selain Excel/CSV.

## 2. Tujuan & Metrik (Sisi Backend)

Tujuan backend:
1. Semua modul menghasilkan jurnal otomatis yang seimbang dan diaudit.
2. API publik + webhook setara fitur internal sejak hari pertama.
3. Keamanan & integritas setara standar sistem keuangan (0 insiden).

| Metrik | Definisi | Target |
|---|---|---|
| Transaksi/bulan | Total dokumen (faktur, jurnal, pembayaran) seluruh tenant | Ditetapkan setelah 3 bulan data awal |
| Insiden keamanan | Kebocoran / akses tidak sah | 0 |
| Keseimbangan jurnal | % jurnal tidak seimbang lolos posting | 0 |
| Nomor ganda | Dokumen bernomor duplikat | 0 |

## 3. Multi-Company, Cabang, Peran & Otorisasi

### 3.1 Struktur organisasi

- Satu akun → banyak perusahaan (multi-company), tiap perusahaan → banyak cabang/gudang.
- Pindah perusahaan cepat via `X-Company-ID` / `X-Branch-ID` header; semua query wajib scoped ke tenant+company+branch.
- Template COA per jenis usaha (retail, jasa, distributor, kontraktor, manufaktur kecil) saat provisioning tenant.

### 3.2 Matriks peran (enforcement server-side)

| Peran | Tugas | Cakupan API |
|---|---|---|
| Owner | Atur perusahaan & pengguna, semua modul | Penuh |
| Akuntan | Jurnal, rekonsiliasi, pajak, tutup buku, laporan | Modul akuntansi + laporan |
| Kasir | Penjualan & terima bayar | Penjualan + kas terbatas |
| Admin gudang | Terima/transfer/opname stok | Persediaan + dokumen terkait |
| Konsultan eksternal | Review/audit, masa berlaku | Baca / terbatas, semua aksi dicatat |

Hak per modul × aksi: `lihat, buat, ubah, hapus, setujui, cetak, ekspor`. Peran dapat dikustom per perusahaan. Setiap request dicek di middleware server — frontend hanya menyembunyikan menu, bukan pengaman.

Konsultan eksternal: token dengan `expires_at`, scope company tertentu, deny-by-default, audit penuh.

## 4. Aturan Umum Semua Modul (Kontrak Backend)

1. **Posting, bukan hapus.** Dokumen `posted` tidak dihapus; hanya `void` atau koreksi via jurnal pembalik. `draft → posted → voided`.
2. **Nomor dokumen:** otomatis, format configurable per tipe (`INV/2026/XI/0001`), unik per company (unique constraint + sequence transaksional, aman konkurensi).
3. **Jurnal seimbang:** `sum(debit) == sum(credit)` per jurnal, divalidasi dalam transaksi DB atomik; tolak jika tidak seimbang.
4. **Kunci periode:** periode tertutup menolak create/update/posting kecuali peran khusus + alasan → audit log.
5. **Audit trail:** setiap mutasi tulis ke tabel log append-only (lihat §7).
6. **Lampiran:** file ke object storage terenkripsi, referensi `attachment_id` di dokumen.
7. **Kurs & pajak:** setiap baris menyimpan `tax_code_id`, `currency`, `exchange_rate`, `rate_date`, `rate_source`.

## 5. Kebutuhan Fungsional per Modul (Aturan Bisnis)

Semua P0 wajib saat rilis. Frontend hanya menampilkan; hitungan resmi di backend.

| Modul | Aturan backend utama |
|---|---|
| Multi-company & cabang | Isolasi `tenant_id, company_id`; konsolidasi antar-cabang via query jurnal terfilter |
| COA | Tree induk-anak, tipe mengikuti Accurate, saldo awal via jurnal pembuka, template per jenis usaha |
| Jurnal umum | Manual, berulang (scheduler worker), penyesuaian, jurnal balik; lampiran wajib opsional |
| Penjualan | Alur quote → order → delivery → invoice → retur → payment; DP dialokasi; diskon & pajak per baris; jurnal: Dr Piutang / Cr Pendapatan, PPN Keluaran |
| Pembelian | Request → order → GRN → invoice → retur → payment; jurnal: Dr Persediaan/Biaya, PPN Masukan / Cr Hutang |
| Piutang & hutang | Kartu pelanggan/pemasok, aging bucket (0-30, 31-60, 61-90, >90), alokasi pembayaran FIFO kecuali override, batas kredit + termin blokir validasi |
| Kas & bank | Kas masuk/keluar, transfer antar kas (jurnal ganda seimbang), rekonsiliasi (match statement vs ledger), impor mutasi CSV |
| Persediaan | Multi-gudang, transfer, opname, adjustment; metode FIFO & rata-rata per item; satuan ganda (konversi), serial/batch tracking |
| Aset tetap | Kategori + umur manfaat; susut otomatis garis lurus & saldo menurun via worker bulanan; pelepasan → jurnal gain/loss |
| Pajak | Kode pajak configurable; PPN masukan/keluaran, PPh 21/23/4(2); bukti potong bernomor; rekap bulanan; tarif sebagai config tanpa redeploy |
| Anggaran | Per akun × periode × departemen × proyek; validasi realisasi vs anggaran |
| Proyek & departemen | Setiap jurnal wajib bisa bawa `project_id, department_id, cost_center_id` nullable untuk P&L segmental |
| Produksi | BoM versi; work order konsumsi bahan (FIFO) + tenaga kerja → barang jadi (harga pokok rata-rata) |
| Penggajian | Karyawan, komponen gaji, tunjangan/potongan, PPh 21, slip; posting jurnal gaji agregat per periode |
| Multi mata uang | Lihat §5.1 |
| Tutup buku | Lihat §5.2 |

### 5.1 Multi Mata Uang & Kurs

- Kurs dari data nyata, bukan dummy. Sumber utama + cadangan (keputusan terbuka: kurs pajak Kemenkeu vs kurs tengah BI vs vendor berbayar). Riwayat kurs disimpan di tabel `exchange_rates`.
- Setiap transaksi asing simpan `currency, rate, rate_date, rate_source`.
- Selisih terealisasi saat pembayaran + belum terealisasi saat revaluasi akhir periode → jurnal otomatis.
- Override manual boleh dengan `override_reason` → audit log.
- API internal: `GET /api/v1/exchange-rates?date=&currency=`; worker sinkron harian.

### 5.2 Tutup Buku & Kunci Periode

- Close bulanan & tahunan: validasi (tidak ada draft menggantung opsional, semua susut & revaluasi posted), lalu jurnal penutup otomatis + set `period_locks`.
- Reopen butuh hak `period:reopen` + alasan; tercatat dan notifikasi owner.
- Laporan komparatif selalu baca dari jurnal posted pada periode terkunci (immutable).

## 6. Laporan (Sumber Data & Kalkulasi)

Semua laporan dihitung dari jurnal posted; tidak ada tabel saldo duplikat sebagai kebenaran (materialized view/cache boleh, tapi rekonsiliasi ke jurnal).

| Kelompok | Laporan | Catatan backend |
|---|---|---|
| Keuangan | Neraca, Laba rugi, Arus kas, Perubahan ekuitas, Neraca saldo, Buku besar, Jurnal umum, Komparatif | Filter `date_from/to, branch_id, department_id, project_id`; arus kas metode tidak langsung dari jurnal |
| AR/AP | Rekap per pelanggan/barang/salesman, aging, kartu, margin | Aging dihitung dari `due_date` vs tanggal laporan |
| Persediaan | Kartu stok, nilai, mutasi/gudang, stok minimum, opname | FIFO layer tracking per gudang |
| Kas & bank | Buku kas/bank, rekonsiliasi, proyeksi arus kas | Proyeksi dari AR/AP jatuh tempo + anggaran |
| Aset | Daftar, jadwal susut, nilai buku | Dari tabel aset + jurnal susut |
| Pajak | Rekap PPN, PPh 21/23/4(2), bukti potong | Format ekspor mengikuti DJP; lihat §7 |
| Proyek/anggaran | P&L per proyek/dept, realisasi vs anggaran | Agregasi dimensi jurnal |
| Produksi & payroll | Biaya/hasil produksi, rekap gaji | Agregat work order & payroll run |
| Kurs | Terealisasi/belum terealisasi, revaluasi | Dari jurnal selisih kurs |
| Audit | Log aktivitas, riwayat dokumen | Query tabel audit append-only |

Export: job async via worker → PDF/Excel di object storage + URL kadaluarsa. Endpoint: `POST /api/v1/reports/{key}/exports`.

## 7. Kepatuhan & Regulasi Indonesia (Sisi Backend)

| Area | Implementasi |
|---|---|
| Standar akuntansi | Template COA & laporan mengikuti Accurate; SAK EMKM/ETAP final dikonfirmasi reviewer (open Q) |
| PPN | `tax_codes` configurable; faktur pajak bernomor; rekap bulanan |
| e-Faktur & Coretax | Ekspor sesuai format DJP; verifikasi ke dok resmi sebelum implementasi (format berubah-ubah) |
| PPh 21/23/4(2) | Engine hitung + bukti potong; tarif di tabel config |
| Kunci periode | §5.2 |
| Retensi | Soft-delete + retensi sesuai ketentuan perusahaan & pajak; purge hanya via job terotorisasi |
| UU PDP | Kolom sensitif (NPWP, rekening, gaji) dienkripsi kolom; akses logged |

**Skema audit log (append-only, tanpa UPDATE/DELETE untuk role biasa):**

```
audit_logs(id, tenant_id, user_id, action, module, document_id,
  old_values JSONB, new_values JSONB, occurred_at, ip, user_agent)
```

Hanya INSERT + SELECT (owner/konsultan berizin). Dapat dicari & diekspor.

## 8. Integrasi, API, Webhook & Migrasi

Prinsip: API publik = API yang dipakai frontend. Desain API dulu.

### 8.1 API publik

- REST + OpenAPI, versioning `/api/v1`, auth API key + OAuth 2.0 (client credentials), rate-limit per tenant, pagination cursor, error Indonesia yang jelas.
- Idempotency-Key untuk POST transaksi agar retry aman.

### 8.2 Webhook

- Event: `invoice.posted, payment.received, stock.moved, document.voided`, dll.
- Setiap endpoint langganan: URL + secret; pengiriman signed HMAC, retry eksponensial, log + dashboard status.
- `POST /api/v1/webhooks/subscriptions`, `GET /api/v1/webhooks/deliveries`.

### 8.3 Konektor pihak ketiga (modular per mitra)

| Mitra | Backend behavior |
|---|---|
| Tokopedia/Shopee | Sinkron order → invoice, stok, pencairan dana → jurnal; worker polling/webhook |
| Payment gateway | Buat VA/QR untuk invoice; callback cocokkan otomatis ke piutang |
| Bank | Impor mutasi CSV; koneksi langsung hanya jika bank sediakan API (konfirmasi per bank) |
| POS | Sinkron penjualan kasir → invoice + kurangi stok |

Catatan: ketersediaan API tiap mitra beda; prioritaskan yang API terbuka (open Q).

### 8.4 Migrasi & ekspor

- Import Excel/CSV dengan template resmi (akun, saldo awal, pelanggan, pemasok, barang, transaksi): upload → preview → validasi per baris → error report → commit atomik.
- Export PDF/Excel untuk laporan & dokumen via job async.

## 9. Non-Fungsional & Keamanan Backend

| Area | Kebutuhan |
|---|---|
| Autentikasi | Kata sandi bcrypt/argon2 + kebijakan; TOTP 2FA; SSO OIDC/SAML; lockout; sesi expiry + revoke |
| Otorisasi | RBAC dicek per request di server |
| Isolasi tenant | DB terpisah per tenant, kredensial terpisah, larangan cross-tenant query; provisioning otomatis (DB + schema + seed) |
| Enkripsi | TLS wajib; at-rest; kolom sensitif enkripsi tambahan; kunci di KMS terpisah dari data |
| Backup | Harian per tenant, kopian off-site terenkripsi, restore drill berkala; RPO/RTO ditetapkan (open Q) |
| AppSec | OWASP Top 10 (SQLi via ORM/param, XSS output-encode, CSRF token, broken access control test), rate-limit, validasi input (zod/validator setara Go), security headers, WAF |
| Secret & supply chain | Secret via env/vault, bukan di kode; scan dependensi + CVE di CI |
| Observabilitas | Log terpusat (tanpa PII berlebih), alert aktivitas curiga, health/uptime monitor |
| Testing | Unit jurnal seimbang + laporan; rekonsiliasi vs Accurate pada data contoh; pentest independen sebelum rilis |
| Performa | Target tps/concurrent ditetapkan (open Q); indeks pada `(tenant_id, company_id, date)`; paginasi |
| Ketersediaan | Target uptime ditetapkan; jendela maintenance diumumkan |
| Lokasi data | Disarankan Indonesia; final di open Q |

Migrasi skema: runner otomatis ke semua DB tenant dengan `schema_version` per tenant; backward-compatible.

## 10. Arsitektur & Tech Stack

```
[Next.js] ──REST──> [Go API] ──> [Postgres per tenant]
                        │  └─> [audit_logs append-only]
                        └─> [Worker Go] ──> laporan/import/backup/kurs/webhook
                        └─> [S3-compatible storage] (lampiran, export PDF/Excel)
```

| Lapisan | Pilihan | Status |
|---|---|---|
| Backend | Go, REST + OpenAPI | Keputusan |
| DB | PostgreSQL (satu DB per tenant) | Rekomendasi |
| Worker | Go (laporan, import, backup, kurs, integrasi) | Rekomendasi |
| File | S3-compatible terenkripsi | Rekomendasi |
| Infra | Docker, CI/CD, server Indonesia | Rekomendasi |

Aturan: nilai uang desimal/integer; jurnal sumber kebenaran; API internal = API publik.

## 11. Roadmap, Risiko & Pertanyaan Terbuka (Backend)

Fase 3 bulanan (usulan, 1 orang — ketat, fitur boleh geser):

- **Fase 1 (bln 1–3):** Tenant provisioning, auth+RBAC, COA, jurnal, audit log, migrasi multi-DB. Gate: fondasi bersih.
- **Fase 2 (bln 4–6):** Penjualan, pembelian, AR/AP, kas-bank, persediaan. Gate: dicoba akuntan sungguhan sebelum lanjut.
- **Fase 3 (bln 7–9):** Aset, pajak, anggaran, proyek, produksi, payroll, kurs, tutup buku, laporan penuh.
- **Fase 4 (bln 10–12):** API publik+webhook, konektor, import/export, hardening, pentest, backup drill. Gate rilis.

| Risiko | Mitigasi backend |
|---|---|
| Scope luas / 1 orang | Library matang, scope gate ketat |
| Kebocoran data | DB per tenant, enkripsi, review tiap fase, pentest |
| Salah hitung | Test jurnal otomatis + banding Accurate + reviewer akuntan |
| Aturan pajak/Coretax berubah | Tarif & format sebagai config |
| API pihak ketiga tidak stabil | Konektor modular, antrian + retry |
| Operasional DB per tenant | Otomasi provisioning & migrasi sejak fase 1 |
| Sumber kurs hilang | Utama + cadangan + cache riwayat |

Pertanyaan terbuka (butuh keputusan sebelum implementasi):
1. Sumber kurs: pajak Kemenkeu / BI / vendor berbayar?
2. Target concurrent & transaksi/bulan (awal & target)?
3. "1.000 pengguna" = terdaftar / aktif / prospek?
4. Cloud & lokasi server, wajib Indonesia?
5. Target uptime, RPO, RTO?
6. SAK tepat (EMKM/ETAP) vs mengikuti Accurate?
7. Prioritas marketplace/bank/gateway pertama?
8. Siapa reviewer akuntansi/pajak selama dev?

---
*Detail layar & alur klik: [PRD Frontend](../frontend/prd.md). Teknis mengikat: [TRD](./trd.md), [ERD](./erd.md), [API](./api/).*
