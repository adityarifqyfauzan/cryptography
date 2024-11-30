package page

import (
	"encoding/hex"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/adityarifqyfauzan/cryptography/crypto"
)

func SHAHash(w fyne.Window) fyne.CanvasObject {
	// Section 1: Hash Data
	inputLabel := widget.NewLabel("Data untuk di-hash:")
	dataInput := widget.NewMultiLineEntry()
	dataInput.SetPlaceHolder("Masukkan data yang akan di-hash")

	hashOutputLabel := widget.NewLabel("Hasil Hash (Hex):")
	hashOutput := widget.NewMultiLineEntry()
	hashOutput.SetPlaceHolder("Hasil hash akan muncul di sini")
	hashOutput.SetMinRowsVisible(3)
	hashOutput.Disable() // Read-only

	hashButton := widget.NewButton("Generate Hash", func() {
		data := dataInput.Text
		if len(data) == 0 {
			dialog.ShowError(errors.New("Data tidak boleh kosong"), w)
			return
		}

		hash := crypto.ManualSHA([]byte(data))
		hashOutput.SetText(hex.EncodeToString(hash))
	})

	copyHashButton := widget.NewButton("Salin Hash", func() {
		if len(hashOutput.Text) == 0 {
			dialog.ShowError(errors.New("Tidak ada hash untuk disalin"), w)
			return
		}
		w.Clipboard().SetContent(hashOutput.Text)
		dialog.ShowInformation("Informasi", "Hash disalin ke clipboard", w)
	})

	resetHashButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm("Konfirmasi", "Apakah Anda yakin ingin mereset data?", func(b bool) {
			if b {
				dataInput.SetText("")
				hashOutput.SetText("")
			}
		}, w)
	})

	// Section 2: Validate Hash
	validateLabel := widget.NewLabel("Validate Hash:")
	validateInput := widget.NewMultiLineEntry()
	validateInput.SetPlaceHolder("Masukkan data asli untuk divalidasi")

	validateHashLabel := widget.NewLabel("Hash untuk Validasi (Hex):")
	validateHashInput := widget.NewMultiLineEntry()
	validateHashInput.SetPlaceHolder("Masukkan hash yang ingin divalidasi")

	validateButton := widget.NewButton("Validate", func() {
		data := validateInput.Text
		hashHex := validateHashInput.Text

		if len(data) == 0 || len(hashHex) == 0 {
			dialog.ShowError(errors.New("Data asli dan hash untuk validasi tidak boleh kosong"), w)
			return
		}

		expectedHash := crypto.ManualSHA([]byte(data))
		expectedHashHex := hex.EncodeToString(expectedHash)

		if expectedHashHex == hashHex {
			dialog.ShowInformation("Validasi Berhasil", "Hash sesuai dengan data asli", w)
		} else {
			dialog.ShowError(errors.New("Hash tidak sesuai dengan data asli"), w)
		}
	})

	resetValidateButton := widget.NewButton("Reset", func() {
		dialog.ShowConfirm("Konfirmasi", "Apakah Anda yakin ingin mereset validasi?", func(b bool) {
			if b {
				validateInput.SetText("")
				validateHashInput.SetText("")
			}
		}, w)
	})

	// Layout untuk Hash Section
	hashSection := container.NewVBox(
		inputLabel,
		dataInput,
		hashButton,
		hashOutputLabel,
		hashOutput,
		container.NewHBox(copyHashButton, resetHashButton), // Tombol berdampingan
	)

	// Layout untuk Validate Hash Section
	validateSection := container.NewVBox(
		validateLabel,
		validateInput,
		validateHashLabel,
		validateHashInput,
		validateButton,
		resetValidateButton,
	)

	// Gabungkan semua section
	content := container.NewVBox(
		hashSection,
		widget.NewSeparator(),
		validateSection,
	)

	return content
}

func SHA(w fyne.Window) fyne.CanvasObject {
	return markdownContent("sha.md")
}
