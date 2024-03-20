package main

import (
	"fmt"

	singleton "github.com/punkestu/singleton-task/singleton"
)

func main() {
	// Coba dapatkan instance singleton
	provider := singleton.GetInstance() // akan membuat instance baru
	provider.AddData("bima")

	fmt.Println()
	// Coba dapatkan instance singleton lagi
	otherProvider := singleton.GetInstance() // akan mengembalikan instance yang sudah ada
	provider.AddData("amib")

	fmt.Println()
	// Coba ambil data dari instance singleton
	data := otherProvider.GetData() // akan mengembalikan semua data yang sudah ditambahkan
	fmt.Println(">> Data yang didapat dari instance singleton")
	for i := range data {
		fmt.Println("+ " + data[i])
	}
}
