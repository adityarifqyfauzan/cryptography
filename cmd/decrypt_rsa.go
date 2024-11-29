package cmd

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/adityarifqyfauzan/cryptography/crypto"
	"github.com/manifoldco/promptui"
)

func DecryptRSA() error {
	prompt := promptui.Prompt{
		Label: "Masukkan Encrypted Data",
	}

	data, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	data = strings.ReplaceAll(data, " ", "")

	prompt = promptui.Prompt{
		Label: "Masukkan Kunci Privat",
	}

	strKey, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	resetTerminal()
	key := strings.Split(strKey, " ")

	d := new(big.Int)
	d, err = crypto.HexToBigInt(key[0])
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	n := new(big.Int)
	n, err = crypto.HexToBigInt(key[1])
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	privateKey := [2]*big.Int{d, n}

	encrypted, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	decrypted, err := crypto.ManualRSADecrypt(privateKey, encrypted)
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}
	resetTerminal()
	fmt.Println("=============== RESULT ===============")
	fmt.Printf("Data Terdekripsi: %s\n", string(decrypted))
	fmt.Println("======================================")
	return nil
}
