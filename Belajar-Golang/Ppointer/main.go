package main

import "fmt"

type Data struct {
	Name, Gender string
}

func main() {
	// // Pass by value
	// di mana ketika kita masukan b = a jika kita ingin mengubah
	// data b maka a tidak ikut berubah karna dia sudah di kopi dan
	// di masukan ke variable lain
	// dataGue := Data{"Farid", "Pria"}
	// dataGueKe2 := dataGue
	// dataGueKe2.Name = "Udin"
	// fmt.Println(dataGue)
	// fmt.Println(dataGueKe2)

	// // Pass by reference dengan Pointer &
	// Kita gunakan & di awal kalimat yang ingin kita copy
	// sehingga kita bisa mengakses datanya dan mengexploitasinya
	// var dataGue Data = Data{"Farid", "Pria"}
	// var dataGueKe2 *Data = &dataGue
	// dataGueKe2.Name = "Udin"
	// fmt.Println(dataGue) // ini ikut berubah
	// fmt.Println(dataGueKe2)

	// // Pass by reference dengan Pointer *
	// Kita gunakan ini di awal nama VARIABLE sehingga dia bisa mengexploitasi
	// semua data yang ada di dalam struct yang sama dan kita tambahkan & di
	// setiap kalimat yang ingin kita ponter
	var addres1 = Data{"Farid", "Pria"}
	var addres4 = &Data{"Farid", "Pria"} //jika kita menaruh & di awal kalimat struct dia akan membuat data baru dan tidak terpengatuh oleh perubahan Pointer *
	var addres2 = &addres1
	var addres3 = addres4 // kta tdk perlu menambah & di awal adres4 di dalam var addres3 karna kita sudah menaruhnya di dalam Data addres4

	// Jika kita ingin mengubah supaya semua data yang ada di dalam struc
	// supaya ikut berubah kita tambahkan * di awal kalimat variable
	// addres2 = Data{"Diana", "Wanita"} // ini tidak ikut berubah jika tidak di beri tanda * dan & di addres2
	*addres2 = Data{"Diana", "Wanita"}
	*addres3 = Data{"Lsi", "llll"} // ini juga tdk perlu & karna ini di dalam addres3 tdk di panggil

	fmt.Println(addres1)
	fmt.Println(addres2)
	fmt.Println(addres3)
	fmt.Println(addres4)

	// // Pass by reference dengan Pointer new
	// kita gunakan pointer new untuk membuat struct menjadi kosong/baru dan bisa
	// mengisi sesuai dengan nama dan type yang ada di dalam struct
	addres5 := new(Data)
	addres5.Name = "Udin"
	addres5.Gender = "Pria lah"
	fmt.Println(addres5)

}
