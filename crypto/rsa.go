package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"math/rand"
)

// ManualRSAEncrypt encrypts the given plaintext using the public key (e, n).
func ManualRSAEncrypt(publicKey [2]*big.Int, plaintext []byte) ([]byte, error) {
	if len(publicKey) != 2 {
		return nil, errors.New("invalid public key length: must be 2 bytes")
	}

	if len(plaintext) == 0 {
		return nil, errors.New("plaintext cannot be empty")
	}

	e := publicKey[0]
	n := publicKey[1]

	// Pad plaintext to fit key size
	paddedPlaintext, err := padDataRSA(plaintext, len(n.Bytes())-1)
	if err != nil {
		return nil, err
	}
	m := new(big.Int).SetBytes(paddedPlaintext)

	// Ensure plaintext is smaller than modulus n
	if m.Cmp(n) >= 0 {
		return nil, errors.New("plaintext too large for key size")
	}

	// Encrypt: c = m^e mod n
	c := new(big.Int).Exp(m, e, n)
	return c.Bytes(), nil
}

// ManualRSADecrypt decrypts the given ciphertext using the private key (d, n).
func ManualRSADecrypt(privateKey [2]*big.Int, ciphertext []byte) ([]byte, error) {
	if len(privateKey) != 2 {
		return nil, errors.New("invalid private key length: must be 2 bytes")
	}

	if len(ciphertext) == 0 {
		return nil, errors.New("ciphertext cannot be empty")
	}

	d := privateKey[0]
	n := privateKey[1]

	c := new(big.Int).SetBytes(ciphertext)

	// Ensure ciphertext is smaller than modulus n
	if c.Cmp(n) >= 0 {
		return nil, errors.New("ciphertext too large for key size")
	}

	// Decrypt: m = c^d mod n
	m := new(big.Int).Exp(c, d, n)
	return unpadDataRSA(m.Bytes()), nil
}

func GenerateRSAKeys(bitSize int) ([2]*big.Int, [2]*big.Int, error) {
	// Generate two random primes p and q
	p := generateRandomPrime(bitSize / 2)
	q := generateRandomPrime(bitSize / 2)
	// Compute n = p * q
	n := new(big.Int).Mul(p, q)

	// Compute phi(n) = (p-1)(q-1)
	phi := new(big.Int).Mul(
		new(big.Int).Sub(p, big.NewInt(1)),
		new(big.Int).Sub(q, big.NewInt(1)),
	)

	// Choose e (public exponent) such that gcd(e, phi) = 1
	e := big.NewInt(65537) // Common choice for e
	if new(big.Int).GCD(nil, nil, e, phi).Cmp(big.NewInt(1)) != 0 {
		return [2]*big.Int{}, [2]*big.Int{}, fmt.Errorf("public exponent e is not relatively prime to phi(n)")
	}

	// Compute d (private exponent) as the modular inverse of e mod phi
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		return [2]*big.Int{}, [2]*big.Int{}, fmt.Errorf("failed to compute private exponent")
	}

	// Return public and private keys
	publicKey := [2]*big.Int{e, n}
	privateKey := [2]*big.Int{d, n}
	return publicKey, privateKey, nil
}

func PublicKeyToBase64(publicKey [2]*big.Int) (string, string) {
	e := base64.StdEncoding.EncodeToString(publicKey[0].Bytes())
	n := base64.StdEncoding.EncodeToString(publicKey[1].Bytes())
	return e, n
}

func PrivateKeyToHex(privateKey [2]*big.Int) (string, string) {
	d := hex.EncodeToString(privateKey[0].Bytes())
	n := hex.EncodeToString(privateKey[1].Bytes())
	return d, n
}

func HexToBigInt(hexStr string) (*big.Int, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(bytes), nil
}

func Base64ToBigInt(base64Str string) (*big.Int, error) {
	bytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(bytes), nil
}

// padDataRSA pads the data to fit within the key size.
func padDataRSA(data []byte, blockSize int) ([]byte, error) {
	if len(data) >= blockSize {
		return nil, fmt.Errorf("plaintext too large for block size")
	}
	padding := blockSize - len(data)
	return append(bytes.Repeat([]byte{0}, padding), data...), nil
}

// unpadDataRSA removes the padding from the decrypted data.
func unpadDataRSA(data []byte) []byte {
	i := 0
	for i < len(data) && data[i] == 0 {
		i++
	}
	return data[i:]
}

// generateRandomPrime generates a random prime number with the specified bit size.
func generateRandomPrime(bitSize int) *big.Int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	upperLimit := new(big.Int).Lsh(big.NewInt(1), uint(bitSize))

	for {
		primeCandidate := new(big.Int).Rand(rng, upperLimit)
		if primeCandidate.ProbablyPrime(20) {
			return primeCandidate
		}
	}
}
