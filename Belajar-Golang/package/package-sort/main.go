package main

import (
	"fmt"
	"sort"
)

// type People struct {
// 	Name, Gender string
// 	Age          int
// }

// type person []People

// Cara ke satu kita gunakan ini
// func (per person) Len() int           { return len(per) }
// func (per person) Less(i, j int) bool { return per[i].Name < per[j].Name }
// func (per person) Swap(i, j int)      { per[i], per[j] = per[j], per[i] }

func main() {
	// people := []People{
	// 	{"Farid anang s", "Pria", 17},
	// 	{"Jodi alfiansyah", "Pria", 21},
	// 	{"Wagas tri sunardi", "Pria", 12},
	// 	{"Raika aisyah", "Wanita", 3},
	// }
	// fmt.Println(people)

	// // sort cara memanggil cara ke 1 seperti ini
	// // sort.Sort(person(people))

	// // cara ke 2 kita gunakan ini yang lebih simple dan ringkas tanpa
	// // mendeklarasikan len less dan swap
	// sort.Slice(people, func(i, j int) bool {
	// 	return people[i].Age < people[j].Age
	// })
	// fmt.Println(people)

	// Search
	angka := []int{1, 2, 10, 4, 5, 6, 7, 5, 3, 5, 4, 3, 2, 4, 5, 6, 2}
	x := 10
	i := sort.Search(len(angka), func(i int) bool { return angka[i] >= x })
	if i < len(angka) && angka[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, angka)
	} else {
		fmt.Printf("%d not found in %v\n", x, angka)
	}
}
