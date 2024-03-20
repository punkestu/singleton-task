package singleton

import "fmt"

// dataProvider dibuat private agar tidak bisa diinisiasi dari luar package
type dataProvider struct {
	data []string
}

// untuk menyimpan instance singleton
var instance *dataProvider

// untuk mendapatkan instance singleton || membuat instance singleton jika belum ada
func GetInstance() *dataProvider {
	fmt.Println(">> Memeriksa instance singleton")
	if instance == nil {
		fmt.Println(">> Membuat instance baru")
		instance = &dataProvider{}
	} else {
		fmt.Println(">> Instance sudah ada, mengembalikan instance yang sudah ada")
	}
	return instance
}

// ==== mulai method dari singleton ====
func (dp *dataProvider) AddData(data string) {
	fmt.Println(">> Menambahkan data ke instance singleton")
	fmt.Println(">> Data yang ditambahkan: " + data)
	dp.data = append(dp.data, data)
}

func (dp *dataProvider) GetData() []string {
	return dp.data
}

// ==== akhir method dari singleton ====
