package main

import "fmt"

func getFullName() (string, int, bool) {
	return "farid", 10, true
}

func main() {
	// firstName, middleName, lastName := getFullName()

	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	// Jika kita hanya ingin menggunakan data pertama saja kita bisa
	// menggunakan _
	// contoh
	// firstName, _, _ := getFullName()
	// fmt.Println(firstName)

	kalimat, number, logika := getFullName()
	fmt.Println(kalimat)
	fmt.Println(number)
	fmt.Println(logika)

}
