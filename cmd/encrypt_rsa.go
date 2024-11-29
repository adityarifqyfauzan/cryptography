package cmd

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/adityarifqyfauzan/cryptography/crypto"
	"github.com/manifoldco/promptui"
)

func EncryptRSA() error {
	prompt := promptui.Select{
		Label: "Pilih",
		Items: []string{
			"Buat Key",
			"Encrypt",
			"◀️  Kembali",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Terjadi kesalahan: %v", err)
	}

	switch result {
	case "Buat Key":
		fmt.Println("Panjang key dalam bit.")
		fmt.Println("Panjang kunci RSA menentukan kekuatan perlindungan terhadap serangan brute force (serangan faktorisasi n).")
		prompt := promptui.Select{
			Label: "Pilih Panjang Key",
			Items: []string{
				"1024",
				"2048",
				"4096",
				"8192",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}
		// string to int
		bitSize, err := strconv.Atoi(result)
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}

		publicKey, privateKey, err := crypto.GenerateRSAKeys(bitSize) // 1024-bit RSA (for simplicity)
		if err != nil {
			return fmt.Errorf("Error generating keys: %v", err)
		}
		resetTerminal()
		publicE, publicN := crypto.PublicKeyToBase64(publicKey)
		privateD, privateN := crypto.PrivateKeyToHex(privateKey)
		fmt.Println("Public Key:", publicE, publicN)
		fmt.Println()
		fmt.Println("Private Key:", privateD, privateN)
		fmt.Println(orange + "Simpan key tersebut dengan baik.")

	case "Encrypt":
		prompt := promptui.Prompt{
			Label: "Masukkan Data",
		}

		data, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}
		plaintext := []byte(data)

		prompt = promptui.Prompt{
			Label: "Masukkan Kunci Publik",
		}

		strKey, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}
		key := strings.Split(strKey, " ")

		e := new(big.Int)
		e, err = crypto.Base64ToBigInt(key[0])
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}

		n := new(big.Int)
		n, err = crypto.Base64ToBigInt(key[1])
		if err != nil {
			return fmt.Errorf("Terjadi kesalahan: %v", err)
		}

		publicKey := [2]*big.Int{e, n}

		// Encrypt
		ciphertext, err := crypto.ManualRSAEncrypt(publicKey, plaintext)
		if err != nil {
			return fmt.Errorf("Error encrypting: %v", err)
		}
		resetTerminal()
		fmt.Printf("Ciphertext (hex): %s\n", hex.EncodeToString(ciphertext))

		return nil
	case "◀️  Kembali":
		return nil
	}

	return nil
}
