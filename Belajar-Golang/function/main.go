package main

import "fmt"

func sayHello(name string, addres string) string {
	if name == "" {
		return "hallo kosong"
	} else {
		return "Hello " + name
	}
}

func main() {
	output := sayHello("", "sesela")
	fmt.Println(output)
	fmt.Println(sayHello("Loli", "sesela"))

}
