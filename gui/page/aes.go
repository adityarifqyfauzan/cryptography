package page

import (
	"encoding/hex"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/adityarifqyfauzan/cryptography/crypto"
)

var (
	aesBitSize = 0
)

func AESDecrypt(w fyne.Window) fyne.CanvasObject {
	// Input untuk pesan terenkripsi
	encryptedMessageLabel := widget.NewLabel("Encrypted Message:")
	encryptedMessageInput := widget.NewMultiLineEntry()
	encryptedMessageInput.SetPlaceHolder("Masukkan pesan terenkripsi (hex string)")

	// Input untuk key
	keyLabel := widget.NewLabel("Key:")
	keyInput := widget.NewEntry()
	keyInput.SetPlaceHolder("Masukkan key (hex string)")

	// Output untuk decrypted message
	decryptedMessageLabel := widget.NewLabel("Decrypted Message:")
	decryptedMessageOutput := widget.NewMultiLineEntry()
	decryptedMessageOutput.SetPlaceHolder("Hasil dekripsi akan muncul di sini")
	decryptedMessageOutput.Disable() // Hanya untuk membaca

	// Tombol untuk melakukan dekripsi
	decryptButton := widget.NewButton("Decrypt", func() {
		encryptedMessage := encryptedMessageInput.Text
		keyHex := keyInput.Text

		// Validasi input
		if len(encryptedMessage) == 0 {
			dialog.ShowError(errors.New("Pesan terenkripsi tidak boleh kosong"), w)
			return
		}
		if len(keyHex) == 0 {
			dialog.ShowError(errors.New("Key tidak boleh kosong"), w)
			return
		}

		// Decode hex string key
		key, err := hex.DecodeString(keyHex)
		if err != nil {
			dialog.ShowError(errors.New("Key tidak valid, harus dalam format hex string"), w)
			return
		}

		// Validasi panjang key
		if len(key) != 16 && len(key) != 24 && len(key) != 32 {
			dialog.ShowError(errors.New("Key harus memiliki panjang 16, 24, atau 32 byte"), w)
			return
		}

		// Decode hex string ciphertext
		ciphertext, err := hex.DecodeString(encryptedMessage)
		if err != nil {
			dialog.ShowError(errors.New("Pesan terenkripsi tidak valid, harus dalam format hex string"), w)
			return
		}

		// Lakukan dekripsi
		plaintext, err := crypto.ManualAESDecrypt(key, ciphertext)
		if err != nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Gagal mendekripsi pesan: %v", err)), w)
			return
		}

		dialog.ShowInformation("Informasi", "Dekripsi berhasil", w)

		// Tampilkan hasil dekripsi
		decryptedMessageOutput.SetText(string(plaintext))
	})

	// Tombol untuk menyalin ciphertext ke clipboard
	copyDecryptedMessageButton := widget.NewButton("Salin Dekripsi", func() {
		if len(decryptedMessageOutput.Text) > 0 {
			w.Clipboard().SetContent(decryptedMessageOutput.Text)
			dialog.ShowInformation("Informasi", "Dekripsi berhasil disalin ke clipboard", w)
		}
	})

	resetButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm(
			"Konfirmasi", "Apakah Anda yakin ingin mereset halaman ini?",
			func(b bool) {
				if b {
					encryptedMessageInput.SetText("")
					keyInput.SetText("")
					decryptedMessageOutput.SetText("")
				}
			}, w)
	})

	// Tata letak halaman
	content := container.NewVBox(
		encryptedMessageLabel,
		encryptedMessageInput,
		keyLabel,
		keyInput,
		decryptButton,
		decryptedMessageLabel,
		decryptedMessageOutput,
		copyDecryptedMessageButton,
		resetButton,
	)

	return content
}

func AESEncrypt(w fyne.Window) fyne.CanvasObject {
	// Pilihan jumlah bit
	bitOptions := []int{128, 192, 256}
	bitSize := widget.NewSelect([]string{"128", "192", "256"}, func(value string) {})
	bitSize.PlaceHolder = "Pilih jumlah bit (Default 128)"

	// Label dan output untuk key
	keyLabel := widget.NewLabel("Key (simpan dengan aman):")
	keyOutput := widget.NewMultiLineEntry()
	keyOutput.SetPlaceHolder("Hasil key akan muncul di sini")
	keyOutput.SetMinRowsVisible(5)
	keyOutput.Disable()

	// Tombol copy key
	copyKey := widget.NewButton("Salin Key", func() {
		if len(keyOutput.Text) > 0 {
			w.Clipboard().SetContent(keyOutput.Text)
			dialog.ShowInformation(
				"Informasi",
				"Simpan key dengan aman untuk proses enkripsi dan dekripsi",
				w,
			)
		}
	})

	// Input untuk pesan
	messageLabel := widget.NewLabel("Pesan untuk dienkripsi:")
	messageInput := widget.NewMultiLineEntry()
	messageInput.SetPlaceHolder("Masukkan pesan yang akan dienkripsi")

	// Output untuk ciphertext
	ciphertextLabel := widget.NewLabel("Ciphertext Output:")
	ciphertextOutput := widget.NewMultiLineEntry()
	ciphertextOutput.SetPlaceHolder("Hasil ciphertext akan muncul di sini")
	ciphertextOutput.SetMinRowsVisible(5)
	ciphertextOutput.Disable()

	// Tombol untuk melakukan enkripsi
	encryptButton := widget.NewButton("Encrypt", func() {
		message := messageInput.Text
		if len(message) == 0 {
			dialog.ShowError(errors.New("Pesan tidak boleh kosong"), w)
			return
		}

		// Decode key dari hex string
		key, err := hex.DecodeString(keyOutput.Text)
		if err != nil {
			dialog.ShowError(errors.New("Format key tidak valid, harus hex string"), w)
			return
		}

		// Ambil nilai bit yang dipilih
		selectedBit := 128 // Default
		for _, option := range bitOptions {
			if bitSize.Selected == fmt.Sprintf("%d", option) {
				selectedBit = option
				break
			}
		}

		if keyOutput.Text == "" || selectedBit != aesBitSize {
			// Generate AES key
			var err error
			key, err = crypto.GenerateAESKey(selectedBit)
			if err != nil {
				dialog.ShowError(errors.New("Gagal membuat key AES"), w)
				return
			}
			keyOutput.SetText(hex.EncodeToString(key))
			aesBitSize = selectedBit
		}

		// Lakukan enkripsi
		ciphertext, err := crypto.ManualAESEncrypt(key, []byte(message))
		if err != nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Gagal mengenkripsi pesan: %v", err)), w)
			return
		}

		dialog.ShowInformation("Informasi", "Pesan berhasil dienkripsi", w)

		// Tampilkan ciphertext
		ciphertextOutput.SetText(hex.EncodeToString(ciphertext))
	})

	// Tombol untuk menyalin ciphertext ke clipboard
	copyCiphertext := widget.NewButton("Salin Ciphertext", func() {
		if len(ciphertextOutput.Text) > 0 {
			w.Clipboard().SetContent(ciphertextOutput.Text)
			dialog.ShowInformation("Informasi", "Ciphertext berhasil disalin ke clipboard", w)
		}
	})

	resetButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm(
			"Konfirmasi", "Apakah Anda yakin ingin mereset halaman ini?", func(b bool) {
				if b {
					messageInput.SetText("")
					keyOutput.SetText("")
					ciphertextOutput.SetText("")
				}
			}, w)
	})

	// Tata letak
	content := container.NewVBox(
		bitSize,
		messageLabel,
		messageInput,
		encryptButton,
		keyLabel,
		keyOutput,
		copyKey,
		ciphertextLabel,
		ciphertextOutput,
		copyCiphertext,
		resetButton,
	)

	return content
}

func AES(w fyne.Window) fyne.CanvasObject {
	return markdownContent("aes.md")
}
