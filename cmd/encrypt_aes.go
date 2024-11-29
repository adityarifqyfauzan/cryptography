package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/adityarifqyfauzan/cryptography/crypto"
	"github.com/manifoldco/promptui"
)

func EncryptAES() error {
	fmt.Println("Panjang key dalam bit.")
	fmt.Println("Panjang kunci AES menentukan kekuatan perlindungan terhadap serangan brute force (serangan faktorisasi n).")
	prompt := promptui.Select{
		Label: "Pilih Panjang Key",
		Items: []string{
			"128",
			"192",
			"256",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	key := []byte{}
	switch result {
	case "128":
		bitSize := 128
		key, err = crypto.GenerateAESKey(bitSize)
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}
		fmt.Println("Key (Hex): ", hex.EncodeToString(key))
	case "192":
		bitSize := 192
		key, err = crypto.GenerateAESKey(bitSize)
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}
		fmt.Println("Key (Hex): ", hex.EncodeToString(key))
	case "256":
		bitSize := 256
		key, err = crypto.GenerateAESKey(bitSize)
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}
		fmt.Println("Key (Hex): ", hex.EncodeToString(key))
	}

	prompt1 := promptui.Prompt{
		Label: "Masukkan Data",
	}

	data, err := prompt1.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	encrypted, err := crypto.ManualAESEncrypt(key, []byte(data))
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	fmt.Printf("Data: %s\n", data)
	fmt.Printf("Encrypted: %x\n", encrypted)

	return nil
}
