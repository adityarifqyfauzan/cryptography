# SHA (Secure Hash Algorithm)

## Apa Itu SHA?
SHA (Secure Hash Algorithm) adalah algoritma hashing yang digunakan untuk menghasilkan sidik jari digital (hash) dari data. Hash ini memiliki beberapa karakteristik penting:
- **Unik**: Setiap data menghasilkan hash yang berbeda. Bahkan perubahan kecil pada data akan menghasilkan hash yang sepenuhnya berbeda.
- **Tidak Bisa Dibalik**: Hash tidak memungkinkan Anda untuk mendapatkan kembali data aslinya. Ini membuatnya ideal untuk keamanan.
- **Panjang Tetap**: Berapa pun ukuran data Anda, hash selalu memiliki panjang yang sama.

SHA banyak digunakan dalam berbagai aplikasi keamanan seperti penyimpanan kata sandi, tanda tangan digital, dan verifikasi integritas data.

---

## Fungsi SHA
Fungsi utama SHA adalah menghasilkan hash unik dari data input. Hash ini adalah representasi tetap dari data asli, sehingga cocok untuk:
- **Keamanan Kata Sandi**: Kata sandi disimpan sebagai hash, sehingga data asli tidak terlihat.
- **Verifikasi Integritas**: Membandingkan hash data saat ini dengan hash data asli memastikan data tidak diubah.
- **Tanda Tangan Digital**: Menjamin bahwa dokumen tidak dimodifikasi sejak ditandatangani.

---

## Cara Kerja SHA
1. **Pengolahan Data**:
   - Data input pertama-tama dipersiapkan (padding) agar panjangnya sesuai dengan aturan algoritma.
   - Padding menambahkan bit tambahan ke data sehingga ukurannya menjadi kelipatan tertentu (misalnya 64 byte).

2. **Proses Blok Data**:
   - Data dipecah menjadi blok-blok kecil dengan panjang tetap.
   - Setiap blok diproses secara independen dengan operasi matematika kompleks untuk menghasilkan nilai sementara.

3. **Penggabungan Nilai**:
   - Nilai sementara dari setiap blok digabung untuk membentuk hash akhir.
   - Proses ini memastikan bahwa setiap perubahan kecil dalam data menghasilkan perubahan besar dalam hash.

4. **Hasil Akhir**:
   - Hash akhir dihasilkan dalam format tertentu, seperti hexadecimal, yang mudah dibaca oleh manusia.

---

## Mengapa SHA Penting?
SHA adalah bagian penting dari keamanan data modern. Karena sifatnya yang unik, tidak bisa dibalik, dan memiliki panjang tetap, SHA digunakan untuk:
- Melindungi data sensitif seperti kata sandi.
- Memastikan data yang dikirimkan atau diterima tidak diubah.
- Memberikan autentikasi pada dokumen atau pesan digital.

---

## Catatan
Meskipun algoritma SHA sederhana dapat digunakan untuk belajar, untuk keamanan tingkat tinggi, disarankan menggunakan implementasi standar seperti `crypto/sha256` di dalam standard library bawaan Go atau alat keamanan modern lainnya.
