package main

import "fmt"

func main() {
	datas := map[string]string{
		"name":   "farid",
		"addres": "lombok",
	}

	fmt.Println(datas)
	fmt.Println(datas["name"])
	fmt.Println(datas["addres"])

}
