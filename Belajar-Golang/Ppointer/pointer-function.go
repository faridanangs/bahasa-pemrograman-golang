package main

import "fmt"

type Data struct {
	Name, Gender string
}

// kita lebih baik menggunakan struct Pointer supaya datanya tidak besar
// karna sering di kopi dan di save di memori yang baru
func changeData(data *Data) {
	// kenapa ini tidak berubah kan kita sudah mengubahnya?
	// karna di sini kita tidak menggunakan pointer
	// cara menggunakan pointer cukup taruh bintang di awal parameter structnya
	// contoh *Data dan taruh & di awal struct Variablenya contoh &dataGue / &Data
	data.Name = "farid"
}
func main() {
	dataGue := Data{"", "pria"}
	changeData(&dataGue)
	fmt.Println(dataGue)
}
