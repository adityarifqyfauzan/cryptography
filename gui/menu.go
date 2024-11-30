package gui

import (
	"fyne.io/fyne/v2"
	"github.com/adityarifqyfauzan/cryptography/gui/page"
)

type Menu struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	Menus = map[string]Menu{
		"welcome": {
			"Welcome",
			"",
			page.Welcome,
		},
		"aes": {
			"AES",
			"",
			page.AES,
		},
		"aes-encrypt": {
			"AES Encrypt",
			"Enkripsi menggunakan AES",
			page.AESEncrypt,
		},
		"aes-decrypt": {
			"AES Decrypt",
			"Dekripsi menggunakan AES",
			page.AESDecrypt,
		},
		"rsa": {
			"RSA",
			"",
			page.RSA,
		},
		"rsa-encrypt": {
			"RSA Encrypt",
			"Enkripsi menggunakan RSA",
			page.RSAEncrypt,
		},
		"rsa-decrypt": {
			"RSA Decrypt",
			"Dekripsi menggunakan RSA",
			page.RSADecrypt,
		},
		"sha": {
			"SHA",
			"",
			page.SHA,
		},
		"sha-hash": {
			"Hash",
			"SHA Hash",
			page.SHAHash,
		},
	}

	MenuIndex = map[string][]string{
		"":    {"welcome", "aes", "rsa", "sha"},
		"aes": {"aes-encrypt", "aes-decrypt"},
		"rsa": {"rsa-encrypt", "rsa-decrypt"},
		"sha": {"sha-hash"},
	}
)
