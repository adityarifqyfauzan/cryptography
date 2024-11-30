
# Cryptography

Cryptography adalah ilmu dan seni mengamankan informasi sehingga hanya pihak yang berwenang yang dapat membaca dan memahaminya. Dalam konteks modern, kriptografi biasanya mengacu pada teknik dan metode matematis untuk melindungi data, baik selama pengiriman (transit) maupun saat disimpan (storage).

---

## 📖 Overview Project

Ini adalah sebuah tugas dari matakuliah **Kriptografi dan Steganografi**, Jurusan **PJJ Tekinik Informatika**, **Universitas Siber Asia** Semester **Ganjil**, Tahun Akademik 2024/2025.

Kelompok: **1** 

Kelas: **IF504**

Repository ini menyediakan fungsi-fungsi untuk:
- Melakukan enkripsi dan dekripsi secara symetric menggunakan algoritma AES.
- Melakukan enkripsi dan dekripsi secara asymetric menggunakan algoritma RSA.
- Melakukan hashing menggunakan algoritma SHA.

Pada project ini hanya mengimplementasikan secara sederhana dari masing masing algoritma diatas, tidak disarankan untuk digunakan di lingkungan production **(For Learning Purpose Only!)**.

---

## 🛠️ Instalasi Bahasa Pemrograman Go

Proyek ini ditulis menggunakan **Go**, jadi pastikan Anda telah menginstal Go sebelum melanjutkan.

### Langkah instalasi Go:
1. Kunjungi halaman [resmi Golang](https://go.dev/dl/).
2. Unduh installer Go yang sesuai dengan sistem operasi Anda (Windows, macOS, atau Linux).
3. Ikuti instruksi instalasi:
   - **Windows**: Jalankan installer `.msi` dan ikuti petunjuknya.
   - **macOS**: Gunakan paket `.pkg` untuk instalasi.
   - **Linux**: Ekstrak file tar.gz dan tambahkan path Go ke variabel `PATH` Anda.
4. Verifikasi instalasi dengan perintah berikut di terminal/command prompt:

   ```bash
   go version
   ```

   Output yang diharapkan:
   ```
   go version go1.xx.x <os/arch>
   ```

---

## 🚀 Install Project as Go Installer Executable

Ikuti langkah-langkah berikut untuk menjalankan proyek ini sebagai executable:

### 1. Clone Repository
Unduh repository ke perangkat lokal Anda dengan perintah berikut:
```bash
git clone https://github.com/adityarifqyfauzan/cryptography.git
cd cryptography
```

### 2. Build Project
Gunakan perintah berikut untuk membuat executable file dari kode sumber:
```bash
go build -o cryptography
```

File executable dengan nama `cryptography` akan dibuat di direktori yang sama.

### 3. Jalankan Program
Eksekusi file hasil build:
```bash
./cryptography
```

### 4. (Opsional) Install Sebagai Go Installer
Untuk membuat program ini dapat diakses dari mana saja di terminal:
```bash
go install
```

Setelah diinstal, Anda dapat menjalankan aplikasi menggunakan nama proyeknya:
```bash
cryptography
```

---

## 📂 Struktur Repository

Berikut adalah struktur umum repository:
```
cryptography/
├── main.go         # File utama untuk menjalankan program
├── cmd/            # Endpoint/Controller masing masing menu
├── crypto/         # Implementasi algoritma kriptografi
├── docs/           # Berisi dokumentasi algoritma yang digunakan
├── gui/            # Implementasi GUI untuk aplikasi Cryptography
├  └── page/        # Page untuk masing masing menu
├── go.mod          # List library/dependensi eksternal yang digunakan
├── go.sum          # Mencatat check sum dari dependensi eksternal yang digunakan
└── README.md       # Dokumentasi proyek
```

---

## 🎯 Kontribusi

Kontribusi terbuka untuk siapa saja yang ingin menambahkan fitur baru atau meningkatkan performa. Silakan buat **Pull Request** atau laporkan masalah melalui **Issues**.

