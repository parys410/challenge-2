package main

import (
	"fmt"
	"strconv"

	"github.com/parys410/warungku"
)

func main() {
	fmt.Printf("Selamat Datang, di Warungku")

	// Create Instance dari Warungku
	myWarung := warungku.Warung{}

	// Inisialisasi Struk Belanja
	myWarung.StrukBelanja = &[]warungku.Barang{}

	// Panggil Interface
	var myWarungIf warungku.WarungInterface = &myWarung

	for {
		fmt.Println("\n\nMenu:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Cetak Struk")

		// Tanya User mau pilih menu apa?
		fmt.Printf("Tolong pilih menu (1, 2, dan 3): ")
		var choice string
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Handle pilihan user
		switch choice {
		case "1":
			var namaInput string
			fmt.Printf("\n\nMasukkan Nama Barang : ")
			_, err := fmt.Scanln(&namaInput)
			if err != nil {
				fmt.Println("User input error")
				return
			}

			var qtyInput string
			fmt.Printf("Masukkan Qty : ")
			_, err = fmt.Scanln(&qtyInput)
			if err != nil {
				fmt.Println("User input error")
				return
			}

			var hargaInput string
			fmt.Printf("Masukkan Harga : ")
			_, err = fmt.Scanln(&hargaInput)
			if err != nil {
				fmt.Println("User input error")
				return
			}

			var barang warungku.Barang

			barang.Nama = namaInput
			barang.Qty,_ = strconv.Atoi(qtyInput)
			barang.Harga,_ = strconv.Atoi(hargaInput)

			myWarungIf.TambahBarang(&barang)
			fmt.Println(barang, "Added to List")

		case "2":
			fmt.Printf("\n* Kurangi Barang *\n")

			var namaInput string
			fmt.Printf("\n\nMasukkan Nama Barang yang dihapus: ")
			_, err := fmt.Scanln(&namaInput)
			if err != nil {
				fmt.Println("User input error")
				return
			}

			var qtyInput string
			fmt.Printf("Masukkan Qty yang dihapus: ")
			_, err = fmt.Scanln(&qtyInput)
			if err != nil {
				fmt.Println("User input error")
				return
			}

			qtyDihapus, _ := strconv.Atoi(qtyInput)
			myWarungIf.KurangiQty(namaInput, qtyDihapus)

		case "3":
			fmt.Printf("\n * Cetak Struk *\n")
			myWarungIf.CetakStruk()
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}