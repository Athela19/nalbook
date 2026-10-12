# TRD Frontend — Nalbook

> Implementasi teknis dari [PRD Frontend](./prd.md). Token visual mengikat ke [Design System](./design.md). Data & aturan: [PRD Backend](../backend/prd.md) + [ERD Backend](../backend/erd.md).

## 1. Stack & Versi Baseline

| Lapisan | Pilihan | Catatan |
|---|---|---|
| Framework | Next.js 14+ (App Router), React 18+, TypeScript `strict` | Server Components untuk daftar/laporan, Client untuk form |
| Styling | CSS variables dari `design.md` + Tailwind (mapping token) atau CSS Modules | Satu sumber token, tanpa warna hardcode |
| Form | `react-hook-form` + `zod` (validasi format klien; kebenaran tetap backend) | Pertahankan nilai saat validasi gagal |
| Data | `fetch` cache + `SWR`/`TanStack Query` untuk server-state; state form lokal | Paginasi cursor backend |
| Tabel | `TanStack Table` + virtualisasi (`@tanstack/virtual`) untuk tabel besar | Angka kanan tabular, teks kiri, divider halus, Debit/Kredit sejajar + total; kolom dapat diatur, density toggle |
| Command palette | `cmdk` untuk `Ctrl+K` | Cari menu/dokumen/aksi |
| Angka/tanggal | `Intl.NumberFormat('id-ID')`, `Intl.DateTimeFormat('id-ID')`, `tabular-nums` | `Rp 1.250.000,00`; tanggal transaksi ≠ jatuh tempo ≠ posting |
| Auth klien | Cookie `httpOnly` (`SameSite=Lax`, `Secure`) + refresh otomatis; guard di `middleware.ts` per peran | Token tidak di `localStorage`; CSRF token |
| Tema | `light` default, `dark` sebagai varian (navy gelap, bukan hitam murni — design §11) | Simpan preferensi via backend/mechanism yang ditetapkan |
| Ikon | Lucide 16/18/20px, outline konsisten | Ikon + label untuk Debit/Kredit/status |
| Test | Vitest (unit) + Playwright (e2e alur kasir→laporan) | Tanpa data dummy sebagai produksi |

Struktur yang disarankan:

```
app/(auth)/login app/(dashboard)/{dashboard,penjualan,pembelian,kas,persediaan,aset,pajak,laporan,pengaturan}/...
components/ui/{Button,Input,CurrencyInput,DatePicker,DataTable,Badge,Dialog,Drawer,Toast,EmptyState,Skeleton}.tsx
lib/{api-client.ts,auth.ts,format-id.ts,errors-id.ts}
styles/tokens.css  (mapping design.md §12)
features/*.md (spesifikasi per fitur)
```

## 2. Token Wajib (dari `design.md`, jangan duplikasi nilai)

`styles/tokens.css` memetakan `design.md` §12 persis (sumber kebenaran tunggal): `--color-brand-primary:#142B4A` / `-hover:#1C3A60` / `-active:#0E2038`, `--color-brand-accent:#16A36A` / `-hover:#128456`, background/surface `#FFFFFF`, `surface-secondary:#F5F7FA`, teks `#172B4D` / `#526176` / `#778397`, border `#E1E7EF` / `-strong:#C6D0DD`, `surface-hover:#F0F4F8`, semantik + surface (`success #15803D/#F0FDF4`, `warning #B45309/#FFFBEB`, `danger #B91C1C/#FEF2F2`, `info #1D4ED8/#EFF6FF`), font `Inter, ui-sans-serif, system-ui, sans-serif` skala `xs 12 → 3xl 30`, radius `4/6/8/12/9999px`, spacing kelipatan 4px. Petakan ke config Tailwind, jangan bikin sistem paralel.

Aturan mengikat:
- Primary button = Deep Navy bg + putih; secondary putih + border netral + teks navy; aksen hijau selektif; danger merah; satu aksi utama per halaman; semua state (hover/focus/active/disabled/loading) + cegah double-submit; destruktif pakai konfirmasi proporsional (design §7).
- Card untuk kelompok info (label, nilai, periode, perbandingan bila ada, sumber bila perlu); tanpa bayangan besar; tanpa % tanpa data.
- Tabel: angka kanan + `tabular-nums`, teks kiri, divider halus, kolom Debit/Kredit sejajar + total mudah ditemukan; header berhierarki; pagination/jumlah/filter/sort konsisten; wajib empty/loading/error/kosong-search state.
- Badge soft-bg + teks jelas (Draft, Menunggu Persetujuan, Posted, Lunas, Belum Dibayar, Terlambat, Dibatalkan); status bayar ≠ status posting; tidak hanya warna.
- Dialog untuk konfirmasi/form pendek/perhatian; drawer untuk detail/form kontekstual; trap fokus, close jelas, cegah hilang data.
- Ikon Lucide outline 16/18/20 (nav 18–20, tombol 16) + label untuk yang ambigu + accessible name.
- Kontras AA (4.5:1 normal / 3:1 besar), fokus keyboard visible, target WCAG, tabel `<th>` benar, `prefers-reduced-motion`.
- Motion: 120–180 ms hover/fokus, 150–220 ms dropdown/drawer, easing konsisten, tanpa loop mengganggu, tetap bisa tanpa animasi.
- Dark mode hanya setelah light stabil; permukaan navy gelap bukan hitam murni; jangan invert mentah; simpan preferensi.
- DoD halaman (`design.md` §14): token konsisten, hierarki jelas, tabel terbaca, semua state ada, bisa keyboard, status eksplisit, responsif utuh, kontras AA, tanpa dummy.

## 3. API Client (kontrak ke backend `/api/v1`)

- Satu `api-client.ts`: `baseURL=process.env.NEXT_PUBLIC_API_URL`, header `X-Company-ID`/`X-Branch-ID` dari switcher, `Idempotency-Key` (uuid per submit) untuk POST transaksi, cegah double-submit (button loading).
- Error backend `{code,message}` → kalimat Indonesia santai via `errors-id.ts` (contoh: `PERIOD_LOCKED` → "Periode sudah dikunci. Minta Akuntan/Owner untuk membuka dengan alasan.").
- Preview jurnal: `GET` kalkulasi backend, tampil read-only dengan label Debit/Kredit eksplisit (jangan hanya warna).
- Export PDF/Excel & import: polling job (`export_jobs`/`import_batches`) → progres → unduh presigned URL; import tampilkan error per baris sebelum commit.
- route `middleware.ts`: cek sesi + peran per perusahaan; sembunyikan menu + blokir route; backend tetap penegak akhir.

## 4. Layar & Performa (PRD §4–§6)

- Shell: sidebar kiri (expanded/collapsed) + header (breadcrumb, `Ctrl+K`, notifikasi, profil); mobile = drawer/bottom-nav; tabel boleh horizontal scroll terstruktur.
- Daftar: pagination server default 25–50, filter/sort di URL (shareable), skeleton bukan spinner penuh, empty-state + CTA.
- Form transaksi: header + garis + ringkasan + preview jurnal + Simpan draf/Posting; shortcut keyboard; CurrencyInput rata kanan; DatePicker ID.
- Dashboard: kartu kas/bank, AR/AP jatuh tempo, P&L berperiode jelas; angka klik → laporan terfilter; dilarang %/grafik tanpa data.
- Cetak: CSS print A4/A5/struk, preview live, logo + kolom configurable.
- Target awal: LCP <2.5s, interaksi tabel <100ms (virtualisasi), debounce search 300ms.

## 5. Aksesibilitas & i18n

- Semua string Indonesia; tidak ada bahasa lain di rilis ini. Pesan validasi menjelaskan cara memperbaiki.
- Label di atas field, placeholder hanya contoh; field wajib ditandai; pertahankan input aman saat gagal.
- Tabel pakai `<th>` benar; status/error tidak via warna saja; dialog trap fokus + cegah hilang data.

## 6. Keterlacakan

| PRD Frontend § | TRD § | Design § |
|---|---|---|
| §3 prinsip/sistem | §2 token + komponen `ui/` | §1–§7, §12–§14 DoD |
| §4 navigasi | §4 shell | §4 layout |
| §5 alur modul | §3 client + §4 form | §8 accounting UX |
| §6 dashboard/laporan/cetak | §4 performa + cetak | §7 tabel/badge, §8 dashboard |
| §7 NFR | §4–§5 | §10 a11y, §9 motion, §11 dark |
