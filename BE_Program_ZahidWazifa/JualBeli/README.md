# Jual dan Beli Barang Bekas

Proyek ini adalah sistem backend untuk aplikasi "Jual dan Beli Barang Bekas". Sistem ini mendukung proses CRUD (Create, Read, Update, Delete) untuk setiap entitas yang terlibat dalam aplikasi.

## Struktur Proyek

- **src/**: Berisi semua kode sumber untuk aplikasi.
  - **BE_ErdEndpoint_ZahidWazifa.go**: File utama yang mengatur logika aplikasi dan server.
  - **controllers/**: Berisi controller untuk menangani logika bisnis.
    - **barang_controller.go**: Mengelola operasi CRUD untuk entitas Barang.
    - **transaksi_controller.go**: Mengelola operasi CRUD untuk entitas Transaksi.
    - **user_controller.go**: Mengelola operasi CRUD untuk entitas User.
  - **models/**: Berisi model untuk berinteraksi dengan database.
    - **barang.go**: Model untuk entitas Barang.
    - **transaksi.go**: Model untuk entitas Transaksi.
    - **user.go**: Model untuk entitas User.
  - **routes/**: Mengatur routing untuk aplikasi.
    - **routes.go**: Menghubungkan endpoint dengan metode controller.
  - **config/**: Berisi konfigurasi database.
    - **database.go**: Mengatur koneksi ke database.
  - **utils/**: Berisi fungsi utilitas.
    - **response.go**: Fungsi untuk memformat respons yang dikirim ke klien.

## Cara Menjalankan Proyek

1. Pastikan Anda memiliki Go terinstal di sistem Anda.
2. Clone repositori ini.
3. Jalankan perintah `go mod tidy` untuk menginstal dependensi.
4. Jalankan aplikasi dengan perintah `go run src/BE_ErdEndpoint_ZahidWazifa.go`.
5. Akses API melalui endpoint yang telah ditentukan.

## Lisensi

Proyek ini dilisensikan di bawah [Lisensi MIT](LICENSE).