package main

import "fmt"

// Type
type Blacklist func(string) bool

func register(name string, black Blacklist) {
	if black(name) {
		fmt.Println("You are blocking", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	register("admin", blacklist)

	register("farid", blacklist)

	register("root", func(name string) bool {
		return name == "root"
	})
}
