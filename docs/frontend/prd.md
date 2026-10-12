# PRD Frontend — Nalbook

> Dokumen ini adalah sisi frontend dari PRD Nalbook. Logika hitung & keamanan ada di [PRD Backend](../backend/prd.md).
> Sumber: PRD Nalbook 12 Okt 2026 oleh Muhammad Syarif Nurrohman. Bahasa UI: Indonesia saja.

## 1. Konteks & Ruang Lingkup Frontend

Nalbook menargetkan UMKM dan kantor akuntan yang butuh fungsi setara Accurate tetapi UI modern, lapang, dan cepat dipakai seharian. Rilis production-ready 12 bulan, tim 1 orang, web responsif (tanpa native).

| Item | Keputusan Frontend |
|---|---|
| Framework | Next.js (React + TypeScript), responsif mobile |
| Bahasa UI | Indonesia saja, format tanggal/angka Indonesia |
| Tema | Terang sebagai baseline; gelap menyusul setelah stabil (navy gelap, bukan hitam murni — `design.md` §11) |
| API | Hanya via backend Go (`/api/v1`); frontend tidak hitung jurnal resmi |
| Offline | Tidak ada mode offline |
| Billing | Tidak ada langganan/paket di rilis ini |

In-scope UI: dashboard, pemilih perusahaan/cabang, form transaksi + preview jurnal, pusat laporan + filter tersimpan, pusat import/ekspor, pengaturan perusahaan/pengguna/peran/keamanan, cetak template kustom.

Out-of-scope: aplikasi native, offline, bahasa selain Indonesia.

## 2. Pengguna, Peran & Implikasi UI

| Segmen | Kebutuhan | Implikasi UI |
|---|---|---|
| UMKM (retail, jasa, distributor, kontraktor, manufaktur kecil) | Catat jual/stok/kas cepat, laporan kilat | Template COA siap pakai, form pendek, default cerdas |
| Kantor akuntan | Kelola banyak klien satu login | Switcher perusahaan cepat, hak per klien, audit log bisa dicari |

| Peran | Yang dilihat di UI |
|---|---|
| Owner | Semua menu + Pengaturan + Pengguna + Audit log |
| Akuntan | Modul akuntansi + Laporan + Pajak + Tutup buku |
| Kasir | Hanya Penjualan + Kas (tombol lain disembunyikan, backend tetap menolak) |
| Admin gudang | Hanya Persediaan + transfer/opname |
| Konsultan eksternal | Mode baca / terbatas + badge masa berlaku, banner "akses audit" |

Aturan: menu disaring per peran per perusahaan; penolakan backend ditampilkan sebagai pesan ramah ("Anda tidak punya akses, hubungi Owner").

## 3. Prinsip Desain & Sistem Desain

Prinsip produk mengikat ke [Design System](./design.md) §1 (Trustworthy, Precise, Modern, Clear, Professional, Efficient):
1. **Clarity over decoration** — informasi > ornamen; bukan landing-page marketing.
2. **Data over visual noise** — tabel/angka/status mudah dipindai (design §2–§3).
3. **Consistency over novelty** — satu pola komponen di semua modul.
4. **Precision by default** — format `Rp 1.250.000,00`, tanggal ID, status eksplisit.
5. **Progressive complexity** — lanjutan tetap ada tanpa bebani pemula.
6. **Feedback for every action** — posting/batal selalu ada umpan balik eksplisit, bukan hanya warna/animasi.
7. **Input cepat** — keyboard penuh, shortcut, autocomplete, pencarian global `Ctrl+K`.
8. **Alur pendek** — faktur → bayar tanpa banyak pindah halaman (drawer/modal + aksi lanjutan).

Sistem desain mengikat ke [Design System](./design.md) dan [TRD Frontend](./trd.md) — jangan hardcode nilai:
- Token warna (`design.md` §2, §12): primer Deep Navy `#142B4A` / hover `#1C3A60` / active `#0E2038`; aksen Emerald `#16A36A` selektif (bukan semua tombol); teks `#172B4D` / `#526176` / `#778397`; border `#E1E7EF` / `#C6D0DD`; semantik `success #15803D`, `warning #B45309`, `danger #B91C1C`, `info #1D4ED8` + surface lembut. Status tidak boleh hanya warna; debit/kredit tidak boleh hanya merah/hijau.
- Tipografi Inter (`Inter, ui-sans-serif, system-ui, sans-serif`), skala `xs 12 / sm 14 / base 16 / lg 18 / xl 20 / 2xl 24 / 3xl 30`, angka finansial `tabular-nums` rata kanan.
- Bentuk & elevasi: radius `4/6/8/12/9999px`; border tipis netral; bayangan minimal (card tanpa/ringan, dropdown ringan, modal jelas; tanpa glass/neumorphism).
- Ikon Lucide outline 16/18/20px (nav 18–20, tombol 16); ikon + label untuk Debit/Kredit/status ambigu.
- Komponen: Button (primary navy, secondary putih-navy, accent hijau selektif, danger merah; satu aksi utama per halaman; semua state + cegah double-submit), Input (label di atas, placeholder contoh saja, wajib ditandai, pertahankan nilai saat gagal), CurrencyInput (`Rp 1.500.000`), DatePicker ID, DataTable, Badge status + teks, Drawer/Modal, Toast, EmptyState, Skeleton.
- Pola form transaksi: header dokumen + tabel baris + ringkasan pajak/diskon + **pratinjau jurnal (read-only dari API)** + tombol Simpan draf / Posting.

Aksesibilitas (`design.md` §10): target WCAG 2.2 AA, kontras 4,5:1 (teks normal) / 3:1 (besar), fokus keyboard visible, semua fungsi utama bisa keyboard, `prefers-reduced-motion` didukung.

## 4. Struktur Layar & Navigasi

Navigasi utama (`design.md` §4): sidebar kiri + header (breadcrumb, pencarian/`Ctrl+K`, notifikasi, profil); sidebar expanded/collapsed; mobile memakai drawer/navigasi sesuai, bukan sekadar mengecilkan desktop.

| Layar | Isi |
|---|---|
| Dashboard | Kartu kas & bank, piutang/hutang jatuh tempo, P&L berjalan, grafik jual/beli, tabel stok menipis. Klik kartu → drill ke laporan. |
| Pemilih perusahaan & cabang | Switcher di header; search perusahaan; badge cabang aktif; ingat pilihan terakhir |
| Transaksi | Daftar dokumen (filter status draf/posted/void, pencarian, aksi massal terbatas) + Form (lihat §5) |
| Laporan | Katalog laporan per kelompok (§6) + filter tanggal/cabang/dept/proyek tersimpan + export |
| Import & Ekspor | Upload Excel/CSV → preview → error per baris → commit; riwayat job export PDF/Excel + unduh |
| Pengaturan | Perusahaan, cabang, pengguna, peran, nomor dokumen, pajak, kurs, periode, keamanan (2FA/SSO/sesi), webhook, API key |

Pencarian global `Ctrl+K`: lompat ke menu, buka faktur/pelanggan/barang, buat baru.

State kosong & loading: empty-state dengan CTA ("Buat faktur pertama"), skeleton bukan spinner penuh.

## 5. Alur Pengguna per Modul (Tanpa Logika Hitung)

Semua hitung di backend; frontend validasi format + tampilkan hasil API + preview jurnal.

| Modul | Alur layar |
|---|---|
| COA | Tree induk-anak, drag/expand, saldo awal wizard, pilih template usaha |
| Jurnal umum | Form debit/kredit dengan balance indicator live; jurnal berulang (jadwal); lampiran drag-drop |
| Penjualan | Quote → Order → Kirim → Faktur → Retur dalam satu thread dokumen (timeline); tombol "Terima pembayaran" langsung dari faktur; diskon & pajak per baris |
| Pembelian | Request → Order → Terima → Faktur → Retur → Bayar; alur sama dengan penjualan |
| Piutang/Hutang | Aging dengan bucket warna, kartu pelanggan/pemasok, alokasi pembayaran otomatis + override, peringatan limit kredit |
| Kas & Bank | Form kas masuk/keluar, transfer antar kas (dua sisi), layar rekonsiliasi split-view (mutasi vs ledger, match/unmatch), impor mutasi |
| Persediaan | Kartu stok per gudang, transfer (sumber→tujuan), opname (selisih → adjustment approval), input serial/batch via scan/ketik |
| Aset tetap | Daftar + detail (nilai buku live), jadwal susut tabel, tombol pelepasan dengan konfirmasi gain/loss |
| Pajak | Form kode pajak dropdown, rekap bulanan tabel, unduh bukti potong, badge "tarif dari config" |
| Anggaran | Input per akun×periode×dept×proyek; bar realisasi vs anggaran dengan progress bar |
| Proyek/Dept | Picker proyek/dept di setiap form; P&L segmental drill-down |
| Produksi | BoM versi list; work order (bahan → hasil) dengan status; biaya tenaga kerja input |
| Penggajian | Tabel karyawan, run gaji per periode, slip preview + cetak, tombol posting jurnal (akuntan saja) |
| Multi mata uang | Currency picker + kurs tampil (nilai, tanggal, sumber); badge override + kolom alasan |
| Tutup buku | Wizard: checklist validasi → preview jurnal penutup → kunci; tombol buka-kembali (hak khusus + alasan + konfirmasi) |

Aturan UI umum: nomor otomatis tampil tapi bisa edit format di Pengaturan; status badge + teks (`Draft, Menunggu Persetujuan, Posted, Lunas, Belum Dibayar, Terlambat, Dibatalkan` — design §7; status bayar ≠ status posting); bedakan tanggal transaksi/jatuh tempo/posting; saldo tidak hanya warna; riwayat perubahan timeline; posted → tombol Batalkan (bukan Hapus) + alasan + konfirmasi proporsional.

## 6. Dashboard, Laporan, Cetak & Ekspor UI

Dashboard (`design.md` §8, tanpa data contoh/dummy):
- Kartu: saldo kas & bank, pendapatan & beban, piutang, utang, ringkasan arus kas — selalu berperiode jelas + sumber data bila perlu; tanpa % pertumbuhan tanpa data.
- Daftar: transaksi terbaru + tindakan perlu perhatian + stok menipis + faktur jatuh tempo.
- Semua angka klik → laporan terfilter; tanpa grafik dekoratif.

Pusat laporan (kelompok sama dengan backend):
Keuangan (neraca, laba rugi, arus kas, perubahan ekuitas, neraca saldo, buku besar, jurnal, komparatif) · Penjualan & piutang · Pembelian & hutang · Persediaan · Kas & bank · Aset · Pajak · Proyek/dept/anggaran · Produksi & gaji · Kurs · Audit.

Fitur: filter tanggal/cabang/dept/proyek, simpan filter favorit, drill dari angka ke jurnal, export PDF/Excel (job async + progres + unduh), print-friendly.

Cetak dokumen (faktur, surat jalan, kwitansi, slip gaji, bukti potong):
- Template kustom: logo, kolom tampil/sembunyi, ukuran kertas (A4/A5/struk), header/footer.
- Preview live + unduh PDF.

## 7. NFR Frontend

| Area | Kebutuhan |
|---|---|
| Responsif | Desktop penuh + tabel lebar penuh; tablet sesuaikan kolom tanpa hilangkan info; HP: fungsi inti (lihat laporan, buat faktur, catat kas, setujui) via drawer/scroll terstruktur, tanpa zoom horizontal (`design.md` §4) |
| Performa & motion | LCP cepat, paginasi server, debounce search, virtualisasi tabel besar, skeleton; transisi 120–180 ms (hover/fokus), 150–220 ms (dropdown/drawer), dukung `prefers-reduced-motion`; target angka di `trd.md` |
| Keamanan klien | Token di httpOnly cookie / memory, auto-refresh + logout saat revoke; tidak simpan PII di localStorage; CSRF token |
| Error & validasi | Validasi inline (format Rp, tanggal ID), error API diterjemahkan ke kalimat santai |
| Tanpa offline | Banner offline "butuh internet", antrean tidak dijanjikan |
| Cetak | CSS print, hindari potong tabel |

## 8. Arsitektur Frontend

```
[Next.js App Router + TS] → [API Client /api/v1 + Auth] → [Go Backend]
  ├─ Server Components untuk daftar/laporan (SEO & cepat)
  ├─ Client Components untuk form interaktif, tabel, preview jurnal
  ├─ State: server-state via fetch-cache/SWR + form-state lokal
  └─ i18n: string Indonesia + `Intl.NumberFormat('id-ID')`, `Intl.DateTimeFormat('id-ID')`
```

- Satu client API terpusat (base URL env, attach company/branch header, refresh token, map error).
- Desain mobile-first untuk 4 layar inti; sisanya adaptif.
- Lihat `trd.md` untuk versi lib & struktur folder; `features/` untuk spesifikasi per fitur (contoh: `features/auth.md`).

## 9. Roadmap, Risiko & Pertanyaan Terbuka (Frontend)

Definisi selesai halaman mengikat ke `design.md` §14: token konsisten, hierarki jelas, tabel/angka terbaca, semua state (loading/empty/error/success) ada, bisa keyboard + validasi jelas, status eksplisit, responsif tanpa hilangkan info, kontras + fokus AA, tanpa dummy sebagai produksi.

Fase (selaras backend, 1 orang):

- **Fase 1:** Sistem desain + auth + shell (sidebar/switcher) + COA + jurnal.
- **Fase 2:** Jual/beli/AR/AP/kas/stok + rekonsiliasi UI; uji ke akuntan sungguhan (gate).
- **Fase 3:** Aset/pajak/anggaran/proyek/produksi/payroll/kurs/tutup buku + laporan penuh + cetak.
- **Fase 4:** Import/ekspor UI, pengaturan webhook/API key, dark-mode polish, hardening, uji HP.

| Risiko | Mitigasi UI |
|---|---|
| Terasa berat seperti pembanding | Tes 5-tugas inti (buat faktur <2 mnt), potong klik |
| Inkonsistensi 1 orang | Sistem desain + komponen dulu, jangan styling per halaman |
| Tabel besar lambat | Paginasi + virtualisasi sejak awal |
| Salah paham pajak | Copy dipandu akuntan reviewer, bukan istilah teknis |

Pertanyaan terbuka relevan frontend:
- Contoh dashboard Accurate yang paling disukai pengguna untuk ditiru?
- Template cetak bawaan apa saja (faktur A4/A5/struk, surat jalan, kwitansi)?
- Prioritas HP: faktur & kas dulu, atau semua modul?

---
*Aturan bisnis & API: [PRD Backend](../backend/prd.md). Teknis mengikat: [TRD](./trd.md), [Design](./design.md), [Fitur](./features/). Skema data: [ERD Backend](../backend/erd.md).*
