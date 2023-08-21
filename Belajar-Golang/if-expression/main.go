package main

import "fmt"

// if expression
// func main() {
// 	name := "farid"

// 	if name == "anangs" {
// 		fmt.Println("hello anangs")
// 	} else if name == "farid" {
// 		fmt.Println("hello farid")
// 	} else {
// 		fmt.Println("hello loli")
// 	}

// 	// short statement
// 	if length := len(name); length > 2 {
// 		fmt.Println("namamu panjang mazzeh")
// 	} else {
// 		fmt.Println("Ok namamu pass")
// 	}
// }

// switch and case expression
func main() {
	name := "udin"

	// switch name {
	// case "farid":
	// 	fmt.Println("hello farid")
	// case "anang":
	// 	fmt.Println("hello anang")
	// default:
	// 	fmt.Println("heelo boleh kenalan")
	// }

	// short statement
	// switch length := len(name); length > 3 {
	// case true:
	// 	fmt.Println("namamu panjang")
	// case false:
	// 	fmt.Println("namau keren and pass")
	// }

	// tanpa expression di switchnya

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("namamu panjang")
	case length > 5:
		fmt.Println("namamu pass")
	default:
		fmt.Println("namamu keren")
	}
}
