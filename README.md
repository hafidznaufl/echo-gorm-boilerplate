# Echo + GORM + GORM MySQL Driver

Projek boilerplate yang saya miliki adalah kerangka awal untuk mengembangkan aplikasi berbasis Go yang kuat dan efisien dengan teknologi-teknologi kunci berikut:

- Echo: Framework web yang ringan dan cepat untuk membangun aplikasi web dan RESTful API dengan Go. Echo memiliki banyak fitur bawaan yang memungkinkan Anda untuk dengan cepat membangun layanan web yang handal dan efisien.

- GORM: ORM (Object-Relational Mapping) yang kuat untuk Go yang memungkinkan Anda - berinteraksi dengan database dengan mudah. GORM memudahkan pemodelan data Anda dalam kode Go, sehingga Anda dapat mengakses database Anda dengan nyaman dan aman.

- GORM Driver MySQL: Driver khusus MySQL untuk GORM. Ini memungkinkan Anda berkomunikasi dengan database MySQL secara langsung dari aplikasi Go Anda menggunakan GORM.

- Air (Hot Reload): Alat pengembangan yang memungkinkan Anda mengaktifkan fitur "hot reload" untuk aplikasi Anda. Ini sangat membantu dalam mengembangkan dan menguji perubahan kode tanpa perlu menghentikan dan memulai ulang server secara manual.

Dengan menggunakan boilerplate ini, Anda dapat memulai proyek Go Anda dengan cepat, mengintegrasikan basis data dengan mudah menggunakan GORM, dan menjalankan server Anda dengan dukungan hot reload untuk pengembangan yang lebih efisien. Ini adalah dasar yang kuat untuk membangun berbagai jenis aplikasi web dan layanan RESTful dengan Go.

## Installation

- Clone a Repository
```bash
git clone https://github.com/hafidznaufl/echo-gorm-boilerplate.git && cd echo-gorm-boilerplate
```

- Get & Install All Dependencies
```bash
go mod tidy
```
## File Environment `.env.example`

File `.env.example` adalah file contoh konfigurasi yang digunakan dalam proyek ini. File ini berisi daftar variabel lingkungan yang harus diatur dalam file `.env` yang sesungguhnya untuk menjalankan proyek dengan benar. Silakan salin file ini sebagai referensi untuk mengatur variabel lingkungan yang sesuai.

### Variabel Lingkungan yang Diperlukan

Berikut adalah daftar variabel lingkungan yang diperlukan dalam file `.env`:

1. **DB_USER**: Nama pengguna database.
2. **DB_PASS**: Kata sandi pengguna database.
3. **DB_HOST**: Host database.
4. **DB_PORT**: Port database.
5. **DB_NAME**: Nama database yang digunakan.

### Cara Menggunakan `.env.example`

- Duplikat file `.env.example` sebagai `.env` dan membuatnya secara otomatis apabila belum tersedia

   ```bash
   cp -n .env.example .env
   ```
- Isi nilai variabel lingkungan pada `.env` dengan lingkungan yg anda miliki

## Run App
```bash
air
```

## List Documentation

Dalam proyek ini, kami menggunakan beberapa teknologi kunci untuk membangun layanan web yang kuat dan efisien. Berikut adalah daftar link ke dokumentasi resmi dan repository GitHub untuk masing-masing teknologi tersebut:

- **Echo**
  - [Dokumentasi Resmi Echo](https://echo.labstack.com/)
  - [Echo GitHub Repository](https://github.com/labstack/echo)

- **GORM**
  - [Dokumentasi Resmi GORM](https://gorm.io/docs/)
  - [GORM GitHub Repository](https://github.com/go-gorm/gorm)

- **GORM Driver untuk MySQL**
  - [Dokumentasi GORM tentang Menghubungkan ke Database MySQL](https://gorm.io/docs/connecting_to_the_database.html#MySQL)

- **Air (Hot Reload)**
  - [Dokumentasi Resmi Air (Hot Reload)](https://github.com/cosmtrek/air)
  - [Air (Hot Reload) GitHub Repository](https://github.com/cosmtrek/air)

Silakan klik tautan-tautan di atas untuk mengakses dokumentasi resmi dan repository GitHub dari masing-masing teknologi. Dokumentasi ini akan memberikan informasi lebih lanjut, petunjuk penggunaan, dan referensi yang dibutuhkan untuk bekerja dengan teknologi-teknologi tersebut dalam proyek Anda.

## Note

Dalam proyek ini, kami menggunakan [GORM](https://gorm.io/) sebagai ORM (Object-Relational Mapping) untuk berinteraksi dengan database. GORM sangat kuat dan fleksibel, dan mendukung berbagai jenis database.

### Mengganti Driver GORM

Jika Anda ingin menggunakan driver GORM yang berbeda, Anda dapat mengunjungi [GitHub GORM](https://github.com/go-gorm/gorm) untuk menemukan berbagai driver yang tersedia. GORM memiliki dukungan untuk berbagai jenis database, termasuk PostgreSQL, SQLite, dan banyak lagi.

Pertama, pastikan Anda mengimpor driver yang sesuai dengan proyek Anda. Selanjutnya, konfigurasikan koneksi database Anda sesuai dengan driver yang Anda pilih.

Contoh:
```go
import "gorm.io/driver/sqlite"

db, err := gorm.Open(sqlite.Open("my.db"), &gorm.Config{})
```

## Contributors 

<div style="display: flex; align-items: center;">
   <a href="https://github.com/hafidznaufl" style="display: flex; align-items: center; flex-direction: row;">
    <img src="https://avatars.githubusercontent.com/hafidznaufl?s=50" alt="Avatar Hafidz Naufal" width="50" height="50" style="border-radius: 20px; margin-right: 10px;">
   </a>
</div>

[Hafidz Naufal](https://github.com/hafidznaufl) - Kontributor utama proyek ini.




