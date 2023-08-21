package main

import "fmt"

func main() {
	var (
		nilai   int = 80
		absensi int = 80
		// lulusUjian   bool = nilai > 85
		// lulusAbsensi bool = absensi > 70
	)

	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	// var lulus bool = lulusUjian && lulusAbsensi
	fmt.Println(nilai > 80 && absensi > 75)
	fmt.Println(nilai > 80 || absensi > 75)
	fmt.Println(nilai != absensi)
}
