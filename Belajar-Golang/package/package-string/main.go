package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains kita gunakan untuk mencari nilai di dalamnya
	fmt.Println(strings.Contains("Farid udin", "eko"))

	// Split kita gunakan untuk mengubah string di dalamnya menjadi slice
	fmt.Println(strings.Split("Farid anangg sam", " "))

	// ToLower kita gunakan untuk mengubah nilaistring menjadi huruf kecil semua
	fmt.Println(strings.ToLower("Farid ANANGS S"))

	// ToUpper kita gunakan untuk mengubah nilaistring menjadi huruf Besar semua
	fmt.Println(strings.ToUpper("farid anang s"))

	// ToTitle kita gunakan untuk mengubah nilaistring menjadi huruf Besar di awal kalimat
	fmt.Println(strings.Title("farid anang s"))

	// Trim di guankan untuk menghapus kalimat di awal dan di akhir kalimat
	// berapapun banyaknya
	fmt.Println(strings.Trim("aa  farid anang s  aa", "a"))

	// Join kita gunakan untuk menaruh kalimat di dalam slice/array/map
	s := []string{"farid", "anang", "samudra"}
	fmt.Println(strings.Join(s, ", "))

	// Repeat kita gunakan untuk mengulang kata sebanyak yang kita mau
	fmt.Println("Ba " + strings.Repeat("Na", 2))

	// Replace sama dengan replaceAll yang berbeda di sini ada int di mana int ini berfungsi intuk
	// menentukan berapa jumlah string yang akan di ganti
	fmt.Println(strings.Replace("oink oink oink", "nk", "kiki", -1))

	// ReplaceAll kita gunakan untuk mengganti nilai yang ada di dalam string dengan nilai
	// yang kita mau
	fmt.Println(strings.ToLower(strings.ReplaceAll("FArid farid Farid", "farid", "udin")))

}
