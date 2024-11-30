package page

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/adityarifqyfauzan/cryptography/crypto"
)

func RSAEncrypt(w fyne.Window) fyne.CanvasObject {
	var publicKey [2]*big.Int
	var keyGenerated bool

	// Section 1: Generate Key Pair
	keyLabel := widget.NewLabel("Generate RSA Key Pair:")
	bitOptions := widget.NewSelect([]string{"1024", "2048", "4096"}, func(value string) {})
	bitOptions.PlaceHolder = "Pilih jumlah bit (Default 2048)"

	publicKeyLabel := widget.NewLabel("Public Key (Base64):")
	publicKeyE := widget.NewEntry()
	publicKeyE.SetPlaceHolder("Eksponen publik akan muncul di sini")
	publicKeyE.Disable()
	copyPublicKeyButton := widget.NewButton("Salin Public Key", func() {
		w.Clipboard().SetContent(fmt.Sprintf("%s", publicKeyE.Text))
		dialog.ShowInformation("Informasi", "Public key disalin ke clipboard", w)
	})

	privateKeyLabel := widget.NewLabel("Private Key (Hex):")
	privateKeyD := widget.NewEntry()
	privateKeyD.SetPlaceHolder("Eksponen privat akan muncul di sini")
	privateKeyD.Disable()
	copyPrivateKeyButton := widget.NewButton("Salin Private Key", func() {
		w.Clipboard().SetContent(fmt.Sprintf("%s", privateKeyD.Text))
		dialog.ShowInformation("Informasi", "Private key disalin ke clipboard", w)
	})

	progress := widget.NewProgressBar()
	progress.Hide()

	var generateKeyButton *widget.Button
	generateKeyButton = widget.NewButton("Generate Key Pair", func() {
		generateKeyButton.Disable() // Disable tombol selama proses
		progress.SetValue(0)        // Reset progress
		progress.Show()             // Tampilkan progress bar

		selectedBit := 2048
		if bitOptions.Selected != "" {
			fmt.Sscanf(bitOptions.Selected, "%d", &selectedBit)
		}

		go func() {
			for i := 1; i <= 10; i++ {
				time.Sleep(100 * time.Millisecond) // Simulasi proses dengan delay
				progress.SetValue(float64(i) / 10) // Perbarui progress
			}

			// Generate key setelah progress selesai
			pubKey, privKey, err := crypto.GenerateRSAKeys(selectedBit)

			// Perbarui UI langsung
			progress.Hide()            // Sembunyikan progress bar
			generateKeyButton.Enable() // Aktifkan tombol setelah selesai

			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			// Simpan kunci yang dihasilkan
			publicKey = pubKey
			keyGenerated = true

			pubE, pubN := crypto.PublicKeyToBase64(pubKey)
			privD, privN := crypto.PrivateKeyToHex(privKey)

			// Tampilkan kunci di UI
			publicKeyE.SetText(fmt.Sprintf("%s %s", pubE, pubN))
			privateKeyD.SetText(fmt.Sprintf("%s %s", privD, privN))
		}()
	})

	// Section 2: Encrypt Data
	encryptLabel := widget.NewLabel("Encrypt Data:")
	messageInput := widget.NewMultiLineEntry()
	messageInput.SetPlaceHolder("Masukkan pesan yang akan dienkripsi")

	encryptedMessageLabel := widget.NewLabel("Encrypted Message (Hex):")
	encryptedMessageOutput := widget.NewMultiLineEntry()
	encryptedMessageOutput.SetPlaceHolder("Hasil enkripsi akan muncul di sini")
	encryptedMessageOutput.SetMinRowsVisible(5)
	encryptedMessageOutput.Disable()
	copyEncryptedMessageButton := widget.NewButton("Salin Encrypted Message", func() {
		w.Clipboard().SetContent(encryptedMessageOutput.Text)
		dialog.ShowInformation("Informasi", "Encrypted message disalin ke clipboard", w)
	})

	encryptButton := widget.NewButton("Encrypt", func() {
		if !keyGenerated {
			dialog.ShowError(errors.New("Silakan generate key terlebih dahulu"), w)
			return
		}

		message := messageInput.Text
		if len(message) == 0 {
			dialog.ShowError(errors.New("Pesan tidak boleh kosong"), w)
			return
		}

		encryptedBytes, err := crypto.ManualRSAEncrypt(publicKey, []byte(message))
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		encryptedMessageOutput.SetText(hex.EncodeToString(encryptedBytes))
	})

	resetButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm("Konfirmasi", "Apakah Anda yakin ingin menghapus data?", func(confirmed bool) {
			if confirmed {
				keyGenerated = false
				bitOptions.SetSelected("")
				publicKeyE.SetText("")
				privateKeyD.SetText("")
				messageInput.SetText("")
				encryptedMessageOutput.SetText("")
			}
		}, w)
	})

	// Layout untuk Section 1: Generate Key Pair
	generateKeySection := container.NewVBox(
		keyLabel,
		bitOptions,
		generateKeyButton,
		progress,
		publicKeyLabel,
		publicKeyE,
		copyPublicKeyButton,
		privateKeyLabel,
		privateKeyD,
		copyPrivateKeyButton,
	)

	// Layout untuk Section 2: Encrypt Data
	encryptSection := container.NewVBox(
		encryptLabel,
		messageInput,
		encryptButton,
		encryptedMessageLabel,
		encryptedMessageOutput,
		copyEncryptedMessageButton,
		resetButton,
	)

	// Gabungkan semua section
	content := container.NewVBox(
		generateKeySection,
		widget.NewSeparator(),
		encryptSection,
	)

	return content
}

func RSADecrypt(w fyne.Window) fyne.CanvasObject {
	// Input untuk private key
	privateKeyLabel := widget.NewLabel("Private Key (Hex):")
	privateKeyInput := widget.NewMultiLineEntry()
	privateKeyInput.SetPlaceHolder("Masukkan private key dalam format Hex (eksponen:d|modulus:n)")

	// Input untuk encrypted data
	encryptedDataLabel := widget.NewLabel("Encrypted Data (Hex):")
	encryptedDataInput := widget.NewMultiLineEntry()
	encryptedDataInput.SetPlaceHolder("Masukkan data terenkripsi dalam format Hex")

	// Output untuk hasil decrypted message
	decryptedMessageLabel := widget.NewLabel("Decrypted Message:")
	decryptedMessageOutput := widget.NewMultiLineEntry()
	decryptedMessageOutput.SetPlaceHolder("Hasil dekripsi akan muncul di sini")
	decryptedMessageOutput.SetMinRowsVisible(5)
	decryptedMessageOutput.Disable() // Read-only

	// Tombol untuk melakukan dekripsi
	decryptButton := widget.NewButton("Decrypt", func() {
		// Validasi input
		privateKeyText := privateKeyInput.Text
		encryptedDataText := encryptedDataInput.Text

		if len(privateKeyText) == 0 {
			dialog.ShowError(errors.New("Private key tidak boleh kosong"), w)
			return
		}
		if len(encryptedDataText) == 0 {
			dialog.ShowError(errors.New("Data terenkripsi tidak boleh kosong"), w)
			return
		}

		// Parse private key dari Hex
		privateKeyParts := strings.Split(privateKeyText, " ")
		if len(privateKeyParts) != 2 {
			dialog.ShowError(errors.New("Private key harus berformat eksponen:d modulus:n"), w)
			return
		}

		d, err := crypto.HexToBigInt(privateKeyParts[0])
		if err != nil {
			dialog.ShowError(errors.New("Eksponen (d) tidak valid"), w)
			return
		}
		n, err := crypto.HexToBigInt(privateKeyParts[1])
		if err != nil {
			dialog.ShowError(errors.New("Modulus (n) tidak valid"), w)
			return
		}

		privateKey := [2]*big.Int{d, n}

		// Parse encrypted data dari Hex
		encryptedDataBytes, err := hex.DecodeString(encryptedDataText)
		if err != nil {
			dialog.ShowError(errors.New("Data terenkripsi tidak valid"), w)
			return
		}

		// Decrypt data
		decryptedBytes, err := crypto.ManualRSADecrypt(privateKey, encryptedDataBytes)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Gagal mendekripsi data: %w", err), w)
			return
		}

		decryptedMessageOutput.SetText(string(decryptedBytes))
	})

	// Tombol untuk menyalin pesan didekripsi
	copyButton := widget.NewButton("Salin Pesan", func() {
		if len(decryptedMessageOutput.Text) == 0 {
			dialog.ShowError(errors.New("Tidak ada pesan yang didekripsi untuk disalin"), w)
			return
		}
		w.Clipboard().SetContent(decryptedMessageOutput.Text)
		dialog.ShowInformation("Informasi", "Pesan didekripsi disalin ke clipboard", w)
	})

	// Tombol reset untuk menghapus semua field
	resetButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm(
			"Konfirmasi", "Apakah Anda yakin ingin mereset semua data?", func(b bool) {
				if b {
					privateKeyInput.SetText("")
					encryptedDataInput.SetText("")
					decryptedMessageOutput.SetText("")
				}
			}, w)
	})

	// Tata letak halaman
	content := container.NewVBox(
		privateKeyLabel,
		privateKeyInput,
		encryptedDataLabel,
		encryptedDataInput,
		decryptButton,
		decryptedMessageLabel,
		decryptedMessageOutput,
		copyButton,
		resetButton,
	)

	return content
}

func RSA(w fyne.Window) fyne.CanvasObject {
	return markdownContent("rsa.md")
}
