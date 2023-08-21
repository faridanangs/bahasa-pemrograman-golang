package main

import (
	"fmt"
	"os"
)

func main() {
	// args := os.Args
	// fmt.Println(" ")
	// fmt.Println(args[0])
	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

	// name, err := os.Hostname()
	// if err == nil {
	// 	fmt.Println("HostName:", name)
	// } else {
	// 	fmt.Println("Error:", err.Error())
	// }

	fmt.Println(os.Getenv("APP_USERNAME"))
	fmt.Println(os.Getenv("APP_PASSWORD"))
}
