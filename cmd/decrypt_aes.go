package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/adityarifqyfauzan/cryptography/crypto"
	"github.com/manifoldco/promptui"
)

func DecryptAES() error {
	prompt := promptui.Prompt{
		Label: "Masukkan Encrypted Data",
	}

	data, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	prompt = promptui.Prompt{
		Label: "Masukkan Kunci",
	}

	strKey, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	resetTerminal()

	key, err := hex.DecodeString(strKey)
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	encrypted, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	decrypted, err := crypto.ManualAESDecrypt(key, encrypted)
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}
	resetTerminal()
	fmt.Println("=============== RESULT ===============")
	fmt.Printf("Data Terdekripsi: %s\n", string(decrypted))
	fmt.Println("======================================")

	return nil
}
