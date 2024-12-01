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
	// Section 1: Generate Key Pair
	keyLabel := widget.NewLabel("Generate RSA Key Pair:")
	bitOptions := widget.NewSelect([]string{"1024", "2048", "4096"}, func(value string) {})
	bitOptions.PlaceHolder = "Pilih jumlah bit (Default 2048)"

	publicKeyLabel := widget.NewLabel("Public Key (Base64):")
	publicKeyE := widget.NewEntry()
	publicKeyE.SetPlaceHolder("Eksponen publik akan muncul di sini")
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
		if publicKeyE.Text == "" {
			dialog.ShowError(errors.New("Public key tidak boleh kosong, silahkan generate terlebih dahulu"), w)
			return
		}

		message := messageInput.Text
		if len(message) == 0 {
			dialog.ShowError(errors.New("Pesan tidak boleh kosong"), w)
			return
		}

		key, err := keyToBigInt(publicKeyE.Text, "base64")
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		encryptedBytes, err := crypto.ManualRSAEncrypt(key, []byte(message))
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		dialog.ShowInformation("Informasi", "Pesan berhasil dienkripsi", w)

		encryptedMessageOutput.SetText(hex.EncodeToString(encryptedBytes))
	})

	resetButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm("Konfirmasi", "Apakah Anda yakin ingin menghapus data?", func(confirmed bool) {
			if confirmed {
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

func keyToBigInt(key, format string) ([2]*big.Int, error) {
	parts := strings.Split(key, " ")
	if len(parts) != 2 {
		return [2]*big.Int{}, errors.New("invalid key format")
	}

	var a, b *big.Int
	var err error
	if format == "hex" {
		a, err = crypto.HexToBigInt(parts[0])
		if err != nil {
			return [2]*big.Int{}, errors.New("invalid key format")
		}

		b, err = crypto.HexToBigInt(parts[1])
		if err != nil {
			return [2]*big.Int{}, errors.New("invalid key format")
		}
	} else if format == "base64" {
		a, err = crypto.Base64ToBigInt(parts[0])
		if err != nil {
			return [2]*big.Int{}, errors.New("invalid key format")
		}

		b, err = crypto.Base64ToBigInt(parts[1])
		if err != nil {
			return [2]*big.Int{}, errors.New("invalid key format")
		}
	}

	return [2]*big.Int{a, b}, nil
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

		key, err := keyToBigInt(privateKeyText, "hex")
		if err != nil {
			dialog.ShowError(errors.New("Private key tidak valid"), w)
			return
		}

		// Parse encrypted data dari Hex
		encryptedDataBytes, err := hex.DecodeString(encryptedDataText)
		if err != nil {
			dialog.ShowError(errors.New("Data terenkripsi tidak valid"), w)
			return
		}

		// Decrypt data
		decryptedBytes, err := crypto.ManualRSADecrypt(key, encryptedDataBytes)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Gagal mendekripsi data: %w", err), w)
			return
		}

		dialog.ShowInformation("Informasi", "Pesan berhasil didekripsi", w)

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
