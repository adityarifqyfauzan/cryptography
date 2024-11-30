package crypto

func ManualSHA(data []byte) []byte {
	if len(data) == 0 {
		return nil
	}

	paddedData := padDataSHA(data)

	// Inisialisasi state awal (seperti SHA-256)
	var h0, h1, h2, h3 uint32 = 0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476

	// Proses setiap blok 64-byte
	for len(paddedData) >= 64 {
		block := paddedData[:64]
		paddedData = paddedData[64:]

		// Proses setiap byte dalam blok
		for i := 0; i < 64; i += 4 {
			word := uint32(block[i])<<24 | uint32(block[i+1])<<16 | uint32(block[i+2])<<8 | uint32(block[i+3])
			temp := h0 + ((h1 & h2) ^ (^h1 & h3)) + word
			h0 = h3
			h3 = h2
			h2 = h1
			h1 = temp
		}
	}

	// Gabungkan hasil akhir
	return []byte{
		byte(h0 >> 24), byte(h0 >> 16), byte(h0 >> 8), byte(h0),
		byte(h1 >> 24), byte(h1 >> 16), byte(h1 >> 8), byte(h1),
		byte(h2 >> 24), byte(h2 >> 16), byte(h2 >> 8), byte(h2),
		byte(h3 >> 24), byte(h3 >> 16), byte(h3 >> 8), byte(h3),
	}
}

func padDataSHA(data []byte) []byte {
	length := len(data)
	padded := append(data, 0x80) // Tambahkan bit '1' (0x80)
	for len(padded)%64 != 56 {   // Isi dengan '0' hingga panjangnya % 64 == 56
		padded = append(padded, 0x00)
	}

	bitLength := uint64(length * 8) // Panjang data dalam bit
	for i := 0; i < 8; i++ {
		padded = append(padded, byte(bitLength>>(56-i*8)))
	}
	return padded
}
