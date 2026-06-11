[English](README.md) | Bahasa Indonesia

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

1. Download versi terbaru
2. Extract archive ke folder yang diinginkan
3. Jalankan aplikasi
4. Browser default akan terbuka secara otomatis

Jika browser tidak terbuka secara otomatis, buka alamat berikut:

```text
http://127.0.0.1:41991
```

Playlist musik Anda siap digunakan.

> Pastikan file musik berada di folder Musik bawaan sistem operasi.

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
git clone <repository-url>
cd <project-name>
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
go run . --server
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
go build -o dist/remusika
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
- [ ] Media Session API
- [ ] Playlist

## Lisensi
