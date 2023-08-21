package main

import "fmt"

func main() {
	var result = 10*10 - 10
	fmt.Println(result)

	var (
		a = 10
		b = 20
		c = a + b
	)
	fmt.Println(c != b)

}
