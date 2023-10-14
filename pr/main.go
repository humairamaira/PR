package main

import (
	"fmt"
)

func main() {
	// Buat variabel tagihan
	fmt.Print("Masukkan tagihan: ")
	var tagihan int
	_, err := fmt.Scanf("%d", &tagihan)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Inisialisasi variabel tip
	var tip float64

	// Cek apakah nilai tagihan berada di antara 50 dan 300
	if tagihan >= 50 && tagihan <= 300 {
		// Jika nilai tagihan berada di antara 50 dan 300, tip adalah 15%
		tip = float64(tagihan) * 0.15
	} else {
		// Jika nilai berbeda, tip adalah 20%
		tip = float64(tagihan) * 0.20
	}

	// Hitung total nilai (tagihan + tip)
	totalNilai := float64(tagihan) + tip

	// Tampilkan hasil di konsol
	fmt.Printf("Tagihannya %d, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, totalNilai)
}
