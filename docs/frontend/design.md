# Nalbook — Design System

> A modern accounting experience built on clarity, precision, and trust.

## 1. Brand Identity

### Brand Overview

**Nalbook** adalah accounting software yang membantu bisnis mengelola transaksi, pembukuan, laporan keuangan, dan operasional finansial melalui satu sistem yang terstruktur.

Desain produk harus mencerminkan kualitas software finansial profesional: informatif, akurat, konsisten, dan mudah digunakan.

### Brand Personality

- **Trustworthy:** Memberikan rasa aman dan kepercayaan dalam pengelolaan data keuangan.
- **Precise:** Menampilkan angka, transaksi, dan informasi secara akurat dan terstruktur.
- **Modern:** Menggunakan antarmuka yang bersih dan mengikuti pola desain SaaS modern.
- **Clear:** Mengurangi kompleksitas akuntansi melalui hierarki informasi yang jelas.
- **Professional:** Cocok digunakan oleh UMKM hingga perusahaan yang memiliki proses keuangan kompleks.
- **Efficient:** Membantu pengguna menyelesaikan pekerjaan dengan langkah yang logis dan minim hambatan.

### Design Principles

1. **Clarity over decoration.** Informasi lebih penting daripada ornamen visual.
2. **Data over visual noise.** Tabel, angka, dan status transaksi harus mudah dipindai.
3. **Consistency over novelty.** Komponen dan interaksi harus konsisten di seluruh modul.
4. **Precision by default.** Format angka, mata uang, tanggal, dan status harus konsisten.
5. **Progressive complexity.** Fitur kompleks tetap tersedia tanpa membebani pengguna baru.
6. **Feedback for every action.** Setiap operasi penting memberikan umpan balik yang jelas.

---

## 2. Color System

Palet warna berasal dari identitas logo Nalbook: Deep Navy sebagai fondasi kepercayaan, Emerald Green sebagai aksen finansial, dan White sebagai ruang visual yang bersih.

### Primary Colors

| Token | Hex | Penggunaan |
|---|---|---|
| `brand-primary` | `#142B4A` | Identitas utama, navigasi, heading penting |
| `brand-primary-hover` | `#1C3A60` | Hover pada elemen utama |
| `brand-primary-active` | `#0E2038` | Active dan pressed state |
| `brand-accent` | `#16A36A` | Aksi positif, indikator pertumbuhan, elemen terpilih tertentu |
| `brand-accent-hover` | `#128456` | Hover pada aksen hijau |
| `background` | `#FFFFFF` | Latar halaman utama |
| `surface` | `#FFFFFF` | Card, dialog, dan panel |
| `surface-secondary` | `#F5F7FA` | Area pendukung dan panel sekunder |

### Neutral Colors

| Token | Hex | Penggunaan |
|---|---|---|
| `text-primary` | `#172B4D` | Teks utama |
| `text-secondary` | `#526176` | Deskripsi dan informasi pendukung |
| `text-muted` | `#778397` | Metadata dan teks nonutama |
| `border-default` | `#E1E7EF` | Border input, tabel, dan card |
| `border-strong` | `#C6D0DD` | Divider dan border yang membutuhkan penekanan |
| `surface-hover` | `#F0F4F8` | Hover pada baris dan item navigasi |
| `overlay` | `#142B4A` dengan opacity | Backdrop dialog |

### Semantic Colors

Warna semantik harus dibedakan dari warna identitas brand.

| Token | Warna | Penggunaan |
|---|---|---|
| `success` | `#15803D` | Operasi berhasil dan status selesai |
| `success-surface` | `#F0FDF4` | Latar pesan sukses |
| `warning` | `#B45309` | Peringatan dan tindakan yang memerlukan perhatian |
| `warning-surface` | `#FFFBEB` | Latar peringatan |
| `danger` | `#B91C1C` | Kesalahan, pembatalan, dan tindakan destruktif |
| `danger-surface` | `#FEF2F2` | Latar error |
| `info` | `#1D4ED8` | Informasi dan notifikasi netral |
| `info-surface` | `#EFF6FF` | Latar informasi |

### Color Usage Rules

- Deep Navy digunakan untuk identitas, struktur navigasi, dan elemen penting.
- Emerald Green digunakan secara selektif, bukan untuk semua tombol atau seluruh elemen aktif.
- Status transaksi tidak boleh bergantung pada warna saja. Sertakan label teks atau ikon.
- Gunakan kombinasi warna dengan kontras yang memadai untuk keterbacaan.
- Hindari gradien dekoratif, neon, dan banyak warna aksen dalam satu layar.
- Merah dan hijau tidak boleh menjadi satu-satunya pembeda debit dan kredit; gunakan label eksplisit.

---

## 3. Typography

### Font Family

Prioritaskan **Inter** sebagai font utama karena keterbacaannya dalam antarmuka bisnis dan data yang padat.

Fallback:

`Inter, ui-sans-serif, system-ui, sans-serif`

### Type Scale

| Token | Ukuran | Weight | Penggunaan |
|---|---:|---:|---|
| `text-xs` | 12px | 400–500 | Metadata dan label sekunder |
| `text-sm` | 14px | 400–500 | Tabel, form, navigasi |
| `text-base` | 16px | 400 | Teks isi |
| `text-lg` | 18px | 500–600 | Subheading |
| `text-xl` | 20px | 600 | Heading card |
| `text-2xl` | 24px | 600 | Judul halaman |
| `text-3xl` | 30px | 600–700 | Ringkasan atau angka utama |

### Typography Rules

- Gunakan heading yang ringkas dan deskriptif.
- Hindari terlalu banyak ukuran font dalam satu layar.
- Gunakan font weight 600 untuk penekanan; hindari bold berlebihan.
- Gunakan angka tabular untuk nominal, saldo, dan angka dalam tabel.
- Terapkan format angka Indonesia secara konsisten pada locale Indonesia, misalnya `Rp 1.250.000,00` apabila dua desimal diperlukan.
- Jangan menyembunyikan informasi penting hanya untuk menghemat ruang.

Untuk angka finansial, gunakan fitur font `font-variant-numeric: tabular-nums` jika tersedia.

---

## 4. Layout and Spacing

### Layout Philosophy

Nalbook menggunakan layout SaaS berbasis dashboard dengan navigasi yang konsisten, area kerja yang luas, dan struktur informasi yang mudah dipindai.

Antarmuka tidak boleh terasa seperti landing page marketing. Prioritaskan efisiensi operasional dan kepadatan informasi yang terkontrol.

### Spacing Scale

Gunakan basis kelipatan 4px.

| Token | Nilai | Penggunaan |
|---|---:|---|
| `space-1` | 4px | Jarak ikon dan teks |
| `space-2` | 8px | Jarak elemen berdekatan |
| `space-3` | 12px | Padding komponen kecil |
| `space-4` | 16px | Padding standar |
| `space-5` | 20px | Jarak antarelemen |
| `space-6` | 24px | Padding card dan section |
| `space-8` | 32px | Jarak antarseksi |
| `space-12` | 48px | Pemisahan area besar |

### Application Shell

- Sidebar navigasi di sisi kiri.
- Header utama berisi breadcrumb, pencarian atau tindakan global, notifikasi, dan profil pengguna sesuai kebutuhan.
- Konten utama berada pada area kerja yang fleksibel.
- Halaman tabel dapat menggunakan lebar penuh area konten.
- Sidebar mendukung keadaan expanded dan collapsed jika diperlukan.
- Navigasi tidak boleh menutupi konten atau mengganggu pekerjaan utama.

### Responsive Behavior

- **Desktop:** Sidebar dan tabel menggunakan ruang yang tersedia secara optimal.
- **Tablet:** Navigasi dan kolom tabel disesuaikan tanpa menghilangkan informasi penting.
- **Mobile:** Sidebar berubah menjadi drawer atau navigasi yang sesuai; tabel dapat menggunakan horizontal scrolling atau tampilan alternatif yang terstruktur.

Jangan sekadar mengecilkan seluruh elemen desktop untuk membuat tampilan mobile.

---

## 5. Shape, Border, and Elevation

### Border Radius

| Token | Nilai | Penggunaan |
|---|---:|---|
| `radius-sm` | 4px | Badge dan elemen kompak |
| `radius-md` | 6px | Input dan tombol |
| `radius-lg` | 8px | Card dan panel |
| `radius-xl` | 12px | Dialog dan panel besar |
| `radius-full` | 9999px | Avatar dan pill badge |

### Border

- Gunakan border tipis dan netral.
- Pisahkan area berdasarkan struktur, bukan dengan garis tebal.
- Hindari border dekoratif berulang pada setiap elemen kecil.
- Tabel menggunakan divider horizontal yang halus dan konsisten.

### Shadows

Bayangan digunakan secara terbatas untuk menunjukkan hierarki lapisan.

- Card biasa: tanpa bayangan atau bayangan sangat ringan.
- Dropdown dan popover: bayangan ringan.
- Modal: bayangan lebih jelas untuk memisahkan lapisan.
- Sidebar dan konten utama tidak memerlukan bayangan besar.

Hindari efek glassmorphism, neumorphism, dan bayangan yang berlebihan.

---

## 6. Iconography

Gunakan ikon berbasis garis dengan ketebalan dan proporsi konsisten. Lucide merupakan pilihan yang sesuai.

Ketentuan:

- Ukuran umum: 16px, 18px, atau 20px.
- Ikon navigasi: 18–20px.
- Ikon tombol: 16px.
- Gunakan ikon bersama label untuk tindakan yang berpotensi ambigu.
- Ikon tanpa label hanya digunakan untuk tindakan yang sudah umum dipahami dan tetap memiliki accessible name.
- Hindari mencampur ikon outline dan filled dengan gaya yang tidak konsisten.

Ikon tidak boleh menggantikan label seperti Debit, Kredit, Belum Dibayar, atau Terlambat ketika label tersebut diperlukan untuk keputusan finansial.

---

## 7. Component Design

### Buttons

**Primary button**
- Background Deep Navy.
- Teks putih.
- Digunakan untuk tindakan utama seperti Simpan, Buat Transaksi, atau Posting Jurnal.

**Secondary button**
- Background putih.
- Border netral.
- Teks Deep Navy.

**Accent button**
- Emerald Green digunakan untuk tindakan positif tertentu jika hierarki tindakan mendukungnya.
- Jangan menggunakan hijau pada semua tombol utama.

**Danger button**
- Warna merah digunakan untuk penghapusan atau tindakan destruktif.

Aturan:
- Setiap halaman sebaiknya memiliki satu tindakan utama yang paling menonjol.
- Tombol harus memiliki hover, focus, active, disabled, dan loading state.
- Saat proses berjalan, cegah pengiriman duplikat.
- Tindakan destruktif harus memiliki konfirmasi yang proporsional terhadap dampaknya.

### Input and Forms

- Label ditempatkan secara konsisten di atas field.
- Placeholder merupakan contoh format atau petunjuk singkat, bukan pengganti label.
- Validasi menampilkan pesan yang menjelaskan cara memperbaiki input.
- Field wajib harus ditandai secara jelas.
- Format angka, tanggal, dan mata uang harus konsisten.
- Field nominal mendukung penyelarasan angka ke kanan.
- Pertahankan nilai input ketika validasi gagal jika aman dilakukan.

### Cards

Card digunakan untuk mengelompokkan informasi yang berkaitan, bukan untuk membungkus setiap elemen.

Card ringkasan keuangan dapat menampilkan:
- Label metrik.
- Nilai utama.
- Periode laporan.
- Perbandingan periode sebelumnya jika datanya tersedia.
- Keterangan sumber data bila diperlukan.

Jangan menampilkan persentase pertumbuhan atau indikator finansial yang tidak didukung data.

### Tables

Tabel merupakan komponen inti Nalbook.

Ketentuan:
- Header kolom memiliki hierarki visual yang jelas.
- Nominal dan angka disejajarkan ke kanan.
- Teks deskriptif disejajarkan ke kiri.
- Gunakan angka tabular untuk nilai finansial.
- Baris mendukung hover dan focus jika dapat dibuka atau dipilih.
- Status transaksi ditampilkan dengan badge dan teks.
- Pagination, jumlah data, filter, dan pengurutan harus konsisten.
- Sediakan empty state, loading state, error state, dan hasil pencarian kosong.
- Jangan menghapus informasi penting hanya untuk membuat tabel terlihat minimalis.

Untuk tabel akuntansi, pertahankan kesejajaran kolom Debit dan Kredit. Total debit dan kredit harus mudah ditemukan.

### Badges and Status

Gunakan badge dengan latar warna lembut dan teks yang jelas.

Contoh status:
- Draft
- Menunggu Persetujuan
- Posted
- Lunas
- Belum Dibayar
- Terlambat
- Dibatalkan

Status harus mencerminkan state bisnis yang benar. Jangan menyamakan status pembayaran dengan status posting jurnal karena keduanya memiliki makna berbeda.

### Dialogs and Drawers

Gunakan dialog untuk:
- Konfirmasi tindakan penting.
- Form pendek.
- Informasi yang membutuhkan perhatian pengguna.

Gunakan drawer untuk detail atau form yang perlu mempertahankan konteks halaman utama.

Dialog harus mendukung fokus keyboard, penutupan yang jelas, dan pencegahan kehilangan perubahan yang belum disimpan.

---

## 8. Accounting-Specific UX

Desain Nalbook harus mengutamakan kejelasan informasi finansial, bukan sekadar estetika dashboard.

### Financial Data

- Nominal harus konsisten dalam format mata uang yang ditentukan perusahaan.
- Tanggal transaksi harus dibedakan dari tanggal jatuh tempo dan tanggal posting.
- Debit dan kredit harus memiliki label eksplisit.
- Saldo tidak boleh bergantung pada warna untuk menunjukkan positif atau negatif.
- Periode laporan harus terlihat jelas.
- Pembulatan angka harus konsisten dan tidak menyembunyikan perbedaan nilai.
- Angka nol, nilai kosong, dan data yang belum tersedia harus dibedakan jika maknanya berbeda.

### Transaction Workflow

Setiap workflow transaksi harus memberikan kejelasan mengenai:
- Status saat ini.
- Tindakan yang tersedia.
- Dampak tindakan terhadap catatan keuangan.
- Apakah perubahan masih dapat dilakukan.
- Siapa yang melakukan tindakan jika informasi audit tersedia.

Tindakan posting atau pembatalan transaksi harus memiliki feedback yang jelas. Jangan menggunakan animasi atau perubahan warna sebagai satu-satunya tanda bahwa transaksi berhasil diproses.

### Dashboard

Dashboard sebaiknya memprioritaskan informasi yang membantu pengambilan keputusan, seperti:
- Saldo kas dan bank.
- Pendapatan dan beban.
- Piutang.
- Utang.
- Ringkasan arus kas.
- Transaksi terbaru.
- Tindakan yang memerlukan perhatian.

Metrik harus mengikuti periode yang jelas. Hindari grafik dekoratif tanpa fungsi analitis dan hindari menampilkan data contoh seolah-olah merupakan data nyata.

---

## 9. Interaction and Motion

Animasi digunakan untuk memperjelas interaksi, bukan sebagai dekorasi.

- Transisi hover dan focus: sekitar 120–180 ms.
- Transisi dropdown atau drawer: sekitar 150–220 ms.
- Gunakan easing yang halus dan konsisten.
- Hindari animasi berulang yang mengganggu pekerjaan administratif.
- Sediakan dukungan `prefers-reduced-motion`.
- Feedback penting seperti kegagalan validasi harus muncul tanpa penundaan yang tidak perlu.

Setiap interaksi harus tetap dapat digunakan tanpa animasi.

---

## 10. Accessibility

Targetkan WCAG 2.2 Level AA sebagai sasaran desain aksesibilitas.

Ketentuan:
- Kontras teks normal minimal 4.5:1.
- Kontras teks besar minimal 3:1.
- Fokus keyboard terlihat dengan jelas.
- Semua fungsi utama dapat digunakan menggunakan keyboard.
- Input memiliki label yang dapat diakses.
- Ikon interaktif memiliki accessible name.
- Status dan error tidak disampaikan melalui warna saja.
- Tabel menggunakan struktur header dan data yang benar.
- Target interaksi mengikuti panduan ukuran target WCAG yang relevan.

Aksesibilitas harus diterapkan pada komponen dan alur kerja, bukan hanya diverifikasi melalui palet warna.

---

## 11. Dark Mode

Dark mode dapat disediakan sebagai tema tambahan setelah tema terang stabil.

Prinsip:
- Gunakan permukaan navy gelap, bukan hitam murni untuk seluruh layar.
- Pertahankan kontras teks yang memadai.
- Sesuaikan border dan semantic colors agar tetap terbaca.
- Jangan sekadar membalik semua warna.
- Pertahankan makna status transaksi dan konsistensi warna brand.
- Simpan preferensi tema sesuai mekanisme aplikasi yang ditetapkan.

Tema terang merupakan baseline utama desain Nalbook.

---

## 12. Design Tokens and Implementation

Implementasikan token desain sebagai sumber kebenaran tunggal. Hindari penggunaan nilai warna, radius, spacing, atau tipografi yang tersebar tanpa standar.

Contoh token CSS:

```css
:root {
  color-scheme: light;

  --color-brand-primary: #142B4A;
  --color-brand-primary-hover: #1C3A60;
  --color-brand-accent: #16A36A;

  --color-background: #FFFFFF;
  --color-surface: #FFFFFF;
  --color-surface-secondary: #F5F7FA;

  --color-text-primary: #172B4D;
  --color-text-secondary: #526176;
  --color-text-muted: #778397;

  --color-border: #E1E7EF;
  --color-border-strong: #C6D0DD;

  --color-success: #15803D;
  --color-warning: #B45309;
  --color-danger: #B91C1C;
  --color-info: #1D4ED8;

  --font-sans: Inter, ui-sans-serif, system-ui, sans-serif;

  --radius-sm: 4px;
  --radius-md: 6px;
  --radius-lg: 8px;
  --radius-xl: 12px;

  --space-1: 4px;
  --space-2: 8px;
  --space-3: 12px;
  --space-4: 16px;
  --space-5: 20px;
  --space-6: 24px;
  --space-8: 32px;
  --space-12: 48px;
}
```

Token di atas merupakan baseline awal. Nilai warna semantic, focus ring, disabled state, dan tema gelap harus ditambahkan serta diuji sebelum digunakan sebagai sistem produksi.

Jika proyek menggunakan framework atau library komponen, petakan token ke konfigurasi yang sudah ada alih-alih membuat sistem styling paralel.

---

## 13. Do and Don't

### Do
- Gunakan layout yang terstruktur dan konsisten.
- Prioritaskan keterbacaan tabel dan angka finansial.
- Gunakan whitespace untuk memperjelas kelompok informasi.
- Pertahankan hierarki tindakan yang jelas.
- Gunakan warna aksen secara selektif.
- Tampilkan feedback yang eksplisit untuk operasi bisnis.
- Pastikan desain tetap efektif pada data yang padat.

### Don't
- Jangan menggunakan terlalu banyak card dengan bayangan.
- Jangan menempatkan terlalu banyak warna aksen pada satu halaman.
- Jangan menggunakan grafik hanya sebagai dekorasi.
- Jangan menyamakan status transaksi yang berbeda.
- Jangan menyembunyikan label penting di balik ikon.
- Jangan mengandalkan warna untuk menyampaikan informasi finansial.
- Jangan mengorbankan aksesibilitas demi tampilan minimalis.
- Jangan mengklaim data, saldo, atau hasil transaksi yang belum diverifikasi.

---

## 14. Definition of Done

Sebuah halaman dianggap sesuai dengan design system Nalbook apabila:

- Identitas warna, tipografi, spacing, dan radius konsisten.
- Hierarki informasi jelas.
- Tabel dan angka finansial mudah dibaca.
- Semua kontrol memiliki state yang sesuai.
- Loading, empty, error, dan success state tersedia sesuai kebutuhan.
- Form dapat digunakan dengan keyboard dan memiliki validasi yang jelas.
- Status bisnis ditampilkan secara eksplisit.
- Tampilan responsif tanpa menghilangkan informasi penting.
- Kontras dan fokus keyboard memenuhi target aksesibilitas.
- Tidak ada nilai dummy yang ditampilkan sebagai data produksi.
- Komponen menggunakan token desain yang konsisten.

**Prinsip akhir:** Nalbook harus terlihat seperti alat kerja finansial yang tepercaya dan presisi. Kualitas desain diukur dari seberapa jelas pengguna memahami data dan seberapa aman mereka menyelesaikan pekerjaan, bukan dari banyaknya ornamen visual.