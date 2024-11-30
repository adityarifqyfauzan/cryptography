package crypto

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
)

// GenerateAESKey generates a random AES key of the specified bit size (128, 192, or 256 bits).
func GenerateAESKey(bitSize int) ([]byte, error) {
	// Validate the bit size
	if bitSize != 128 && bitSize != 192 && bitSize != 256 {
		return nil, errors.New("invalid bit size: must be 128, 192, or 256")
	}

	// Calculate the key size in bytes
	byteSize := bitSize / 8

	// Generate random bytes for the key
	key := make([]byte, byteSize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random key: %w", err)
	}

	return key, nil
}

func ManualAESEncrypt(key, plaintext []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, errors.New("plaintext cannot be empty")
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid key length: must be 16, 24, or 32 bytes")
	}

	// Pad plaintext to be a multiple of the key length
	plaintext = padDataAES(plaintext, len(key))

	// Encrypt block by block
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i%len(key)]
	}

	return ciphertext, nil
}

func ManualAESDecrypt(key, ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, errors.New("ciphertext cannot be empty")
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid key length: must be 16, 24, or 32 bytes")
	}

	// Decrypt block by block
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		plaintext[i] = ciphertext[i] ^ key[i%len(key)]
	}

	// Remove padding
	plaintext = unpadDataAES(plaintext)

	return plaintext, nil
}

func unpadDataAES(data []byte) []byte {
	if len(data) == 0 {
		return nil
	}
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

func padDataAES(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, pad...)
}
