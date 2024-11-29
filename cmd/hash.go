package cmd

import (
	"fmt"

	"github.com/adityarifqyfauzan/cryptography/crypto"
	"github.com/manifoldco/promptui"
)

func Hash() error {
	prompt := promptui.Select{
		Label: "Pilih",
		Items: []string{
			"Buat Hash",
			"Validasi Hash",
			"◀️  Kembali",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	switch result {
	case "Buat Hash":
		prompt := promptui.Prompt{
			Label: "Masukkan Data",
		}

		data, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}

		hash := crypto.ManualSHA([]byte(data))
		fmt.Printf("Data: %s\n", data)
		fmt.Printf("Hash: %x\n", hash)

	case "Validasi Hash":
		prompt := promptui.Prompt{
			Label: "Masukkan Data",
		}

		data, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}

		prompt = promptui.Prompt{
			Label: "Masukkan Hash",
		}

		hashed, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}

		hash := crypto.ManualSHA([]byte(data))
		resetTerminal()
		if hashed == fmt.Sprintf("%x", hash) {
			fmt.Printf("Data: %s\n", data)
			fmt.Printf("Hash: %x\n", hash)
			fmt.Println("Hash Valid")
		} else {
			fmt.Printf("Data: %s\n", data)
			fmt.Printf("Hash: %x\n", hash)
			fmt.Println("Hash Tidak Valid")
		}

	case "◀️  Kembali":
		return nil
	}

	return nil
}
