package main

import "fmt"

func getFullName() (fName, mName, lName string) {
	fName = "farid"
	mName = "anang"
	lName = "samudra"
	return

}

func main() {
	f, m, l := getFullName()

	fmt.Println(f)
	fmt.Println(m)
	fmt.Println(l)
}
