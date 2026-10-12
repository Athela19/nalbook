# Deployment — Nalbook (Docker Compose, Production)

File compose: `docker-compose.yml`. Contoh env: `.env.example` (salin ke `.env`, jangan commit).

## Service

| Service | Fungsi | Port host | Wajib sehat sebelum lanjut |
|---|---|---|---|
| postgres | PostgreSQL 16, data di volume `pgdata` | tidak ada | `pg_isready` |
| s3 | LocalStack S3 di volume `s3-data` | tidak ada | `/_localstack/health` |
| s3-init | Buat bucket `S3_BUCKET` sekali jalan (`scripts/s3-init.sh`) | — | exit 0 |
| migrate | Migrasi skema semua tenant, sekali jalan | — | exit 0 |
| backend | API Go `:8080` | `8080` | `/healthz` |
| worker | Job latar (`/app/worker`) | — | proses jalan |
| frontend | Next.js standalone `:3000` | `3000` | `/` HTTP 200 |

Urutan start: postgres + s3 → s3-init + migrate → backend + worker → frontend.

## Environment variable

| Nama | Wajib | Default | Keterangan |
|---|---|---|---|
| POSTGRES_USER / POSTGRES_PASSWORD / POSTGRES_DB | ya | — | Hanya dipakai saat volume `pgdata` masih kosong. Ganti password setelahnya tidak memutar kredensial berjalan, harus via SQL + restart. |
| DATABASE_ADMIN_URL / TENANT_DSN_TPL | ya | — | Koneksi admin + pola DSN per tenant. Boot gagal (exit 1) bila kosong. |
| JWT_SECRET | ya | — | Boot gagal bila kosong. |
| S3_ENDPOINT | ya | — | URL internal, mis. `http://s3:4566`. |
| S3_BUCKET | tidak | `nalbook` | Nama bucket, dipakai `s3-init`, `backend`, `worker`. |
| S3_REGION | tidak | `us-east-1` | Region untuk AWS CLI/SDK. |
| FRONTEND_ORIGIN | tidak | `http://localhost:3000` | Origin untuk CORS backend. |
| NEXT_PUBLIC_API_URL | ya (build) | — | Tertanam ke bundle frontend saat `docker build`. Ganti nilainya wajib rebuild image frontend. |
| PORT | tidak | `8080` | Port listen backend di container. |

## Backup dan restore (rekomendasi, belum otomatis)

Backup database per tenant:

```
docker compose exec postgres pg_dump -U "$POSTGRES_USER" -d "$POSTGRES_DB" -Fc -f /tmp/admin.dump
docker compose cp postgres:/tmp/admin.dump ./backup-admin.dump
```

Restore ke database uji (jangan ke production tanpa change control):

```
docker compose cp ./backup-admin.dump postgres:/tmp/restore.dump
docker compose exec postgres pg_restore -U "$POSTGRES_USER" -d <db_uji> /tmp/restore.dump
```

Objek S3 (lampiran, hasil export) belum ada backup otomatis. Untuk production gunakan object storage berversi di luar Compose.

## Batasan yang diketahui

- `/healthz` backend adalah liveness, belum memeriksa koneksi database/S3 (belum ada driver DB di scaffold).
- Worker belum punya health probe; status `Up` hanya berarti proses hidup.
- `NEXT_PUBLIC_API_URL=http://localhost:8080` benar untuk browser, tetapi fetch sisi server di container frontend tidak menjangkaunya; halaman status menampilkan fallback sampai ada pemisahan URL internal/publik (fase 1).
- Persistensi data S3 LocalStack edisi community tidak dijamin across restart; verifikasi berkala dengan data uji.
- Limit CPU/memory belum ditetapkan; tentukan via load test sebelum production.
- TLS dan firewall di luar Compose; pasang reverse proxy bila diekspos ke jaringan.
