package main

import "fmt"

func main() {
	// coun := 1

	// for coun <= 10 {
	// 	fmt.Println("Masih benar", coun)
	// 	coun++
	// }

	// for statement
	// for coun := 1; coun <= 10; coun++ {
	// 	fmt.Println("Hello", coun)
	// }

	// names := []string{"farid", "anangs", "wagas", "raika"}

	// data := make(map[int]string)

	// // for i := 1; i < len(names); i++ {
	// // 	fmt.Println(names[i])
	// // }

	// // for _, value := range names {
	// // 	// fmt.Println("No", i, value)

	// // 	// ?kita gunakan _ unutk memberitahu go bahwa kita
	// // 	// ?tidak membutuhkan variable tersebut
	// // 	fmt.Println(value)

	// // }

	// for count := 1; count <= 10; count++ {
	// 	data[count] = "hello"
	// }
	// for key, value := range data {
	// 	fmt.Println(key, ";", value)
	// }

	// break dan continue
	// for i := 1; i < 10; i++ {
	// 	// break akan berhenti jika kondisi true dan
	// 	// di dalamnya ada break tdk peduli dengan data
	// 	// yang ada di bawahnya
	// 	// if i == 5 {
	// 	// 	break
	// 	// }
	// 	// fmt.Println("ini print ke ", i)
	// 	// contnu akan di jalankan jika kondisi id terpenuhi
	// 	// dan tdk akan menjalankan data yang ada di dalamnya
	// 	// dan tetap akan melanjutkan program di bawahya
	// 	// if i%2 == 0 {
	// 	// 	continue
	// 	// }
	// 	// fmt.Println("ini print ke ", i)
	// }

	var user string

	fmt.Print("Lovyu gak? ")
	fmt.Scanln(&user)

	if user == "y" {
		fmt.Scanln()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				continue
			}
			if i == 3 {
				fmt.Println("hellllo ", i)
				break
			}
			fmt.Println("Loveyou ke ", i, "diana")
		}
	}

}
