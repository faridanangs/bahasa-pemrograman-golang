package main

import "fmt"

func main() {
	// var names [3]string
	// names[0] = "farid"
	// names[1] = "anang"
	// names[2] = "s"
	// names[2] = "samudra"

	// fmt.Println(names[0], names[1], names[2])

	// Kita juga bisa menggunakan cara seerti ini
	var values = [3]int{
		10,
		20,
		30,
	}
	fmt.Println(values)

	var names = [4]string{
		"anangs anana aaasasas",
		"farid",
		"wagas",
		"raika",
	}
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(len(names[0]), len(values))

}
