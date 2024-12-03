package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/adityarifqyfauzan/cryptography/cmd"
	"github.com/adityarifqyfauzan/cryptography/gui"
	"github.com/manifoldco/promptui"
)

func resetTerminal() {
	fmt.Fprint(os.Stdout, "\033[H\033[2J")
}

func main() {
	// Colors
	reset := "\033[0m"

	purple := "\033[1;35m"
	orange := "\033[1;33m"
	white := "\033[1;37m"

	banner := `
	 ██████╗██████╗ ██╗   ██╗██████╗ ████████╗ ██████╗  ██████╗ ██████╗  █████╗ ██████╗ ██╗  ██╗██╗   ██╗
	██╔════╝██╔══██╗╚██╗ ██╔╝██╔══██╗╚══██╔══╝██╔═══██╗██╔════╝ ██╔══██╗██╔══██╗██╔══██╗██║  ██║╚██╗ ██╔╝
	██║     ██████╔╝ ╚████╔╝ ██████╔╝   ██║   ██║   ██║██║  ███╗██████╔╝███████║██████╔╝███████║ ╚████╔╝ 
	██║     ██╔══██╗  ╚██╔╝  ██╔═══╝    ██║   ██║   ██║██║   ██║██╔══██╗██╔══██║██╔═══╝ ██╔══██║  ╚██╔╝  
	╚██████╗██║  ██║   ██║   ██║        ██║   ╚██████╔╝╚██████╔╝██║  ██║██║  ██║██║     ██║  ██║   ██║   
	 ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝        ╚═╝    ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝  ╚═╝   ╚═╝   
																										 `
	fmt.Println(purple + banner + reset)

	fmt.Println(orange + " Tugas Mata kuliah Kriptografi dan Steganografi " + reset)
	fmt.Println(white + " Kelompok: " + reset + orange + "1" + reset)
	fmt.Println(white + " Kode Kelas: " + reset + orange + "IF504" + reset)
	fmt.Println(white + " Dosen Pengampu: " + reset + orange + "Bpk Abdul Azzam Ajhari, S.Kom., M.Kom" + reset)
	fmt.Println()
	fmt.Println(white + " Tugas membuat aplikasi untuk enkripsi dan dekripsi menggunakan: ")
	fmt.Println("- Symetric (AES) dan Asymetric (RSA) ")
	fmt.Println("- Hash menggunakan SHA")
	fmt.Println()
	fmt.Println(purple + "©" + time.Now().Format("2006") + " Kelompok 1. Powered by Go." + reset)

	// initialize gui app
	var apps = app.NewWithID("com.kelompok-1.cryptography")

	for {
		prompt := promptui.Select{
			Label: "Main Menu",
			Items: []string{
				"🖥️   Buka GUI",
				"🔐  Encrypt",
				"🔓  Decrypt",
				"🔑  Hash",
				"🚪  Keluar",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Println("Terjadi kesalahan:", err)
			continue
		}

		switch result {
		case "🖥️   Buka GUI":
			window := apps.NewWindow("Cryptography")
			gui.Run(apps, window)
			window.ShowAndRun()
			window.Close()
			apps.Quit()
			resetTerminal()
			return

		case "🔐  Encrypt":
			prompt := promptui.Select{
				Label: "Pilih Algoritma",
				Items: []string{
					"AES",
					"RSA",
					"◀️  Kembali",
				},
			}

			_, algo, err := prompt.Run()
			if err != nil {
				resetTerminal()
				fmt.Println("Terjadi kesalahan:", err)
				continue
			}

			switch algo {
			case "AES":
				if err := cmd.EncryptAES(); err != nil {
					resetTerminal()
					fmt.Println(err)
					continue
				}

			case "RSA":
				if err := cmd.EncryptRSA(); err != nil {
					resetTerminal()
					fmt.Println(err)
					continue
				}

			case "◀️  Kembali":
				resetTerminal()
				continue
			}

		case "🔓  Decrypt":
			prompt := promptui.Select{
				Label: "Pilih Algoritma",
				Items: []string{
					"AES",
					"RSA",
					"◀️  Kembali",
				},
			}

			_, algo, err := prompt.Run()
			if err != nil {
				fmt.Println("Terjadi kesalahan:", err)
				continue
			}

			switch algo {
			case "AES":
				if err := cmd.DecryptAES(); err != nil {
					resetTerminal()
					fmt.Println(err)
					continue
				}

			case "RSA":
				if err := cmd.DecryptRSA(); err != nil {
					resetTerminal()
					fmt.Println(err)
					continue
				}

			case "◀️  Kembali":
				resetTerminal()
				continue
			}

		case "🔑  Hash":
			if err := cmd.Hash(); err != nil {
				fmt.Println(err)
				continue
			}

		case "🚪  Keluar":
			resetTerminal()
			fmt.Println("Terima kasih!")
			return
		}
	}

}
