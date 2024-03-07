package main

import (
	"fmt"
	"time"
)

func main() {
	// Mengambil tanggal lahir dari pengguna
	fmt.Print("Masukkan tanggal lahir (YYYY-MM-DD): ")
	var input string
	fmt.Scanln(&input)

	// Mengonversi string input menjadi tipe data time.Time
	tanggalLahir, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Gunakan format YYYY-MM-DD.")
		return
	}

	// Menghitung usia
	usia := hitungUsia(tanggalLahir)

	// Menampilkan hasil
	fmt.Printf("Usia Anda: %d tahun\n", usia)
}

func hitungUsia(tanggalLahir time.Time) int {
	// Mendapatkan waktu saat ini
	sekarang := time.Now()

	// Menghitung selisih tahun
	usia := sekarang.Year() - tanggalLahir.Year()

	// Memeriksa apakah ulang tahun sudah terjadi atau belum
	if sekarang.YearDay() < tanggalLahir.YearDay() {
		usia--
	}

	return usia
}
