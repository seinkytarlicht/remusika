<br />
<div align="center">

<img src="assets/images/logo.png" width="140" alt="Rem - ReMusika">

<br />
<br />

[![Release version](https://img.shields.io/github/v/release/seinkytarlicht/remusika?color=brightgreen&label=Latest&style=for-the-badge)](https://github.com/seinkytarlicht/remusika/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/seinkytarlicht/remusika?style=for-the-badge)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/m/seinkytarlicht/remusika?color=brown&label=Commits&style=for-the-badge)](https://github.com/seinkytarlicht/remusika/commits)
![GitHub Repo stars](https://img.shields.io/github/stars/seinkytarlicht/remusika?color=orange&style=for-the-badge)

[English](README.md) | Indonesia

</div>

# ReMusika

Pemutar musik lokal berbasis Go dan Nuxt yang berjalan melalui browser.

Seluruh proses berjalan secara lokal di perangkat Anda tanpa memerlukan koneksi internet, layanan pihak ketiga, maupun instalasi tambahan.

Cukup jalankan aplikasi, buka browser, dan nikmati koleksi musik Anda.

## Preview

> Screenshot dan demo GIF akan ditambahkan pada versi berikutnya.

## Mengapa Dibuat?

Project ini dibuat sebagai eksperimen untuk membuat aplikasi pemutar musik lokal menggunakan teknologi web dengan pengalaman penggunaan yang tetap sederhana.

Tujuan utama project ini:

- Berjalan sepenuhnya di perangkat pengguna
- Tidak memerlukan koneksi internet
- Tidak memerlukan instalasi Node.js atau runtime tambahan
- Mudah didistribusikan dalam bentuk executable

## Fitur

- Memindai folder Musik secara otomatis
- Streaming audio secara lokal
- Kontrol Play, Pause, Seek, dan Loop
- Integrasi system tray
- Otomatis terbuka di browser default
- Dapat digunakan secara offline

### Menu Tray

- **Show** — Membuka aplikasi di browser
- **Exit** — Menghentikan seluruh service dan menutup aplikasi

## Cara Menggunakan

### Instalasi

1. Download versi terbaru, [di sini](https://github.com/seinkytarlicht/remusika/releases)
2. Jalankan aplikasi
3. Browser default akan terbuka secara otomatis

Jika browser tidak terbuka secara otomatis, buka alamat berikut:

```text
http://127.0.0.1:41991
```

Playlist musik Anda siap digunakan.

> Pastikan file musik berada di folder Musik bawaan sistem operasi.

> Masalah yang ditemukan: Jika halaman menampilkan halaman kosong, coba Hard Reload (`Shift` + `R`) atau coba bersihkan cache browser

## Lokasi Folder Musik

### Windows

```text
C:\Users\<your-username>\Music
```

### Linux

```text
/home/<your-username>/Music
```

## Tech Stack

- Go
- Fiber
- Nuxt

## Menjalankan Project

### Clone Repository

```bash
git clone https://github.com/seinkytarlicht/remusika.git
cd remusika
```

### Menjalankan Frontend

```bash
cd frontend
bun install
bun run dev
```

### Menjalankan Backend

Jalankan dari root project:

```bash
go run cmd/app/main.go --server
```

## Struktur Folder

```text
/
├── config
├── frontend
├── helper
├── internal
└── main.go
```

## Build Aplikasi

### Generate Frontend

```bash
cd frontend
bun run generate
```

### Build Backend

Jalankan dari root project:

```bash
go build -o dist/remusika cmd/app/main.go
```

Hasil build akan tersedia di:

```text
dist/remusika
```

## Roadmap

- [x] Scan library musik
- [x] Audio streaming
- [x] Integrasi tray
- [x] Pencarian lagu
- [ ] Logo
- [x] Media Session API
- [x] Playlist

## Lisensi
