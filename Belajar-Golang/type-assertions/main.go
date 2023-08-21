package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	var result interface{} = random()
	// resultToString := result.(string)
	// fmt.Println(resultToString)

	// lebih baik kita menggunakan
	// switch dan case lebih aman
	switch value := result.(type) {
	case string:
		fmt.Println(value, "Value ini bertipe string")
	case int:
		fmt.Println(value, "Value ini bertipe int")
	case bool:
		fmt.Println(value, "Value ini bertipe boolean")
	default:
		fmt.Println("Type tidak di ketahui")
	}
}
