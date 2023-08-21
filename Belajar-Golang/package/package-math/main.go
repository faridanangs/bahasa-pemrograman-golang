package main

import (
	"fmt"
	"math"
)

func main() {
	// untuk mengabil nilai paling tinggi di antara keduanya
	fmt.Println(math.Max(10, 1))

	// untuk mengabil nilai rendah tinggi di antara keduanya
	fmt.Println(math.Min(10, 1))

	// untuk membulatkan ke atas
	fmt.Println(math.Ceil(5.4))

	// untuk membulatkan ke bawah
	fmt.Println(math.Floor(3.6))

	// ini akan membulatkan ke atas sesuai dengan kondisinya
	// jika 1.4 <= 1.5 = 1
	// jika 1.6 >= 1.5 = 2
	// jika 1.5 >= 1.5 = 2
	fmt.Println(math.Round(7.3))
	fmt.Println(math.Round(7.6))
	fmt.Println(math.Round(7.5))

}
