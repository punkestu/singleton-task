# SINGLETON
Repo ini berisi kode untuk tugas mata kuliah Pola-pola perancangan dimana aku ditugaskan untuk membuat implementasi pola Singleton dalam bentuk kode. Pada repo ini aku menggunakan bahasa Go.

## Apa itu singleton?
Dengan menggunakan Singleton, kita bisa membuat sebuah class atau struct yang hanya akan dibuat sekali tapi tetap bisa digunakan dimana saja. Pola ini cocok digunakan jika kita memang hanya butuh 1 instance saja namun tetap ingin digunakan di banyak tempat. Misal ketika kita ingin melakukan transaksi data ke database atau sumber data lain, maka kita tidak perlu membuat instance baru untuk setiap transaksi namun hanya perlu membuat 1 instance yang nantinya bisa digunakan oleh semua transaksi.

## Singleton + Go
Untuk mengimplementasikan Singleton pada bahasa Go, cara yang paling umum adalah dengan membuat struct dan instance singletonnya menjadi private sehingga tidak memungkinkan untuk dirubah secara manual jika sudah diinisiasi. Dengan begitu untuk melakukan inisiasi kita bisa membuat sebuah fungsi public yang akan membuat instance baru jika belum ada instance yang dibuat lalu mengembalikan instance tersebut sehingga bisa digunakan diluar package.