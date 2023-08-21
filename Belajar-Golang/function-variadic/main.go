package main

import "fmt"

func sumAll(numberss ...string) []string {
	kalimat := make([]string, len(numberss), len(numberss))
	for i, value := range numberss {
		kalimat[i] = value
	}

	return kalimat
}
func main() {
	// fmt.Println(sumAll(98, 89, 89, 9, 89080))

	// Kita juga bisa memasukan slice ke dalam function variadic
	slice := []string{"farid", "anangs", "samudra"}
	total := sumAll(slice...)
	fmt.Println(total)
}
