# AES (Advanced Encryption Standard)

## Apa Itu AES?
AES (Advanced Encryption Standard) adalah algoritma enkripsi simetris yang digunakan secara luas untuk melindungi data sensitif. "Simetris" berarti algoritma ini menggunakan kunci yang sama untuk **enkripsi** dan **dekripsi** data. AES dikenal karena keamanannya yang tinggi dan kecepatannya dalam memproses data.

---

## Fungsi AES
AES digunakan untuk:
- **Melindungi Data Sensitif**: Digunakan dalam aplikasi seperti penyimpanan data terenkripsi, komunikasi terenkripsi, dan database aman.
- **Keamanan Transaksi Online**: Melindungi data transaksi keuangan dan informasi pribadi.
- **Keamanan Jaringan**: Digunakan dalam protokol keamanan jaringan seperti HTTPS dan VPN.

---

## Cara Kerja AES
1. **Kunci Enkripsi**:
   - Kunci AES memiliki panjang tetap: 128-bit, 192-bit, atau 256-bit.
   - Kunci dihasilkan secara acak menggunakan generator bilangan acak yang aman.

2. **Enkripsi**:
   - Data asli (plaintext) dienkripsi menggunakan kunci AES.
   - Jika ukuran data tidak sesuai dengan panjang blok (ukuran kunci), data tersebut diproses dengan **padding** agar sesuai.

3. **Proses Enkripsi Blok demi Blok**:
   - Data diproses dalam blok kecil dengan ukuran sesuai kunci (16, 24, atau 32 byte).
   - Setiap blok data dienkripsi menggunakan operasi XOR sederhana dalam implementasi manual.

4. **Hasil Akhir**:
   - Data terenkripsi (ciphertext) dihasilkan. Ciphertext ini tidak dapat dibaca tanpa kunci enkripsi yang benar.

5. **Dekripsi**:
   - Proses dekripsi menggunakan kunci yang sama dengan enkripsi.
   - Data terenkripsi diproses blok demi blok menggunakan operasi yang sama, menghasilkan data asli.

6. **Padding dan Unpadding**:
   - Padding digunakan untuk memastikan ukuran data sesuai panjang blok.
   - Setelah dekripsi, padding dihapus untuk mendapatkan data asli.

---

## Mengapa AES Penting?
1. **Keamanan Tinggi**:
   - AES adalah algoritma yang sangat sulit diretas dengan teknologi saat ini.
2. **Efisiensi**:
   - AES dapat mengenkripsi dan mendekripsi data dengan cepat, bahkan untuk dataset besar.
3. **Standar Internasional**:
   - Digunakan di berbagai aplikasi, mulai dari perangkat konsumen hingga sistem militer.

---

## Catatan
Implementasi manual seperti di project ini adalah contoh sederhana untuk pembelajaran. Untuk kebutuhan keamanan tingkat tinggi, gunakan standard library seperti **`crypto/aes`** di Go.

