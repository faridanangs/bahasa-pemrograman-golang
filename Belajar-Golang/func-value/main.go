package main

import "fmt"

// function as value
// func getGoodby(name string) string {
// 	return "Hello " + name
// }

// type
type Filter func(string) string

// function as parameter
func sayHello(name string, filter Filter) {
	outputFilter := filter(name)
	fmt.Println("Hello " + outputFilter)
}

func filterSpam(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// function as value
	// goodBy := getGoodby
	// result := goodBy("farid")
	// fmt.Println(result)

	sayHello("farid", filterSpam)
	sayHello("anjing", filterSpam)

}
