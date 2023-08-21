package main

import "fmt"

func main() {
	// var names = [...]string{
	// 	"farid",
	// 	"anangs",
	// 	"samudra",
	// 	"diana",
	// 	"wagas",
	// 	"jodi",
	// 	"raika",
	// }
	// var slice1 = names[3:6]
	// fmt.Println(slice1)
	// fmt.Println(cap(slice1))
	// fmt.Println(len(slice1))

	// jika saya mengubah data yang ada di dalam slice
	// data yang ada di dalam names juga ikut brubah
	// slice1[0] = "Lolipop"

	// fmt.Println(slice1)
	// fmt.Println(names)

	// slice2 := names[4:6]
	// fmt.Println(slice2)
	// slice3 := append(slice2, "append")
	// fmt.Println(slice3)
	// slice3[1] = "change_rakika"
	// fmt.Println(slice3)

	// fmt.Println(slice2)
	// fmt.Println(names)

	newSlice := make([]string, 3, 5)
	newSlice[0] = "farid"
	newSlice[1] = "wagas"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// ini array [...]type_data atau [jumlah]type_data
	iniArray := [...]int{1, 2, 3, 4, 5}

	// slice []type_data
	iniSlice := []int{1, 2, 3, 4, 5}

	// ini terlihat seperti tdk ada perbedaan sebenarnya ini sangat berbeda
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
