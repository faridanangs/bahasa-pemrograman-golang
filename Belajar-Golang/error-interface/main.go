package main

import (
	"errors"
	"fmt"
)

func Pembagi(value, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Tidaak boleh di bagi 0")
	} else {
		result := value / pembagi
		return result, nil
	}
}

func main() {
	// var contohErr = errors.New("Ups error")
	// fmt.Println(contohErr)

	value, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("Nilainya adalah", value)
	} else {
		// untuk memanggiil parameter err kita
		// harus mendampinginya dengan Error()
		// karna ini adalah interface{}
		fmt.Println("Error", err.Error())
	}
}
