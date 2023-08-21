package main

import "fmt"

func main() {
	datas := map[string]string{
		"name":   "farid",
		"addres": "lombok",
	}
	datas["title"] = "siswa"
	fmt.Println(len(datas))
	fmt.Println(datas)
	fmt.Println(datas["name"])
	fmt.Println(datas["addres"])

	delete(datas, "name")

	fmt.Println(datas["name"])
	fmt.Println(datas)
	fmt.Println(len(datas))

}
