# RSA (Rivest–Shamir–Adleman)

## Apa Itu RSA?
RSA adalah algoritma enkripsi asimetris yang digunakan secara luas untuk mengamankan data. Berbeda dengan algoritma simetris seperti AES, RSA menggunakan **dua kunci berbeda**:
1. **Public Key (Kunci Publik)**: Digunakan untuk mengenkripsi data.
2. **Private Key (Kunci Privat)**: Digunakan untuk mendekripsi data.

RSA dirancang untuk memastikan keamanan data selama transmisi, bahkan di lingkungan yang tidak aman.

---

## Fungsi RSA
RSA digunakan dalam berbagai aplikasi keamanan, seperti:
- **Tanda Tangan Digital**: Memastikan bahwa dokumen atau pesan tidak diubah sejak ditandatangani.
- **Keamanan Data**: Melindungi data sensitif selama transmisi.
- **Enkripsi Kunci Simetris**: Mengamankan kunci untuk algoritma enkripsi simetris seperti AES.
- **Protokol Keamanan**: Digunakan dalam protokol seperti HTTPS, TLS, dan VPN.

---

## Cara Kerja RSA

### 1. **Pembuatan Kunci**
RSA menghasilkan pasangan kunci **publik** dan **privat** berdasarkan konsep matematika berikut:
- **Bilangan Prima Besar**: Dua bilangan prima besar dipilih secara acak.
- **Modulus (n)**: Hasil perkalian kedua bilangan prima.
- **Eksponen Publik (e)**: Nilai yang umum digunakan adalah 65537, dipilih karena efisien secara komputasi.
- **Eksponen Privat (d)**: Dihitung sebagai invers modular dari e terhadap fungsi totient Euler dari n.

Kunci publik terdiri dari **(e, n)**, sementara kunci privat terdiri dari **(d, n)**.

### 2. **Enkripsi**
- Data asli (plaintext) dienkripsi menggunakan kunci publik.
- Setiap blok data diubah menjadi bilangan bulat (integer) dan diproses dengan rumus:
  ```
  ciphertext = plaintext^e mod n
  ```
- Hasilnya adalah data terenkripsi (ciphertext), yang hanya dapat didekripsi dengan kunci privat.

### 3. **Dekripsi**
- Ciphertext didekripsi menggunakan kunci privat.
- Proses dekripsi dilakukan dengan rumus:
  ```
  plaintext = ciphertext^d mod n
  ```
- Hasilnya adalah data asli (plaintext).

### 4. **Padding**
- Data yang tidak sesuai dengan ukuran blok RSA diproses dengan **padding** agar sesuai. Padding memastikan keamanan tambahan dengan mencegah serangan tertentu.

---

## Mengapa RSA Penting?
1. **Keamanan Tinggi**:
   - RSA menggunakan matematika kompleks yang membuatnya sulit untuk diretas, bahkan dengan komputer modern.

2. **Digunakan Secara Luas**:
   - RSA adalah algoritma standar dalam banyak aplikasi, mulai dari browser web hingga perangkat keras keamanan.

3. **Tanda Tangan Digital**:
   - RSA memungkinkan autentikasi dokumen dengan memastikan bahwa hanya pemilik kunci privat yang dapat menandatanganinya.

---

## Catatan
RSA sangat kuat, tetapi implementasi manual di project ini hanya cocok untuk pembelajaran. Untuk kebutuhan keamanan tingkat tinggi, gunakan standard library seperti **`crypto/rsa`** di Go.
