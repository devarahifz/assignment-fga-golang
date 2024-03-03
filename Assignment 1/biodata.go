package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getTemanByAbsen(absen int) *Teman {
	temanKelas := map[int]Teman{
		1: {"John Doe", "Jl. Menteng No. 123", "Software Engineer", "Ingin belajar bahasa Go"},
		2: {"Jane Smith", "Jl. Sudirman No. 456", "Data Scientist", "Tertarik dengan performa Go dalam pengolahan data besar"},
		3: {"Alice Johnson", "Jl. Thamrin No. 789", "Web Developer", "Mendengar tentang kemudahan pengembangan aplikasi dengan Go"},
	}

	teman, found := temanKelas[absen]
	if !found {
		return nil
	}

	return &teman
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Gunakan: go run biodata.go <nomor_absen>")
		return
	}

	absen := args[1]
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	teman := getTemanByAbsen(absenInt)
	if teman == nil {
		fmt.Println("Tidak ada teman dengan nomor absen tersebut")
		return
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
