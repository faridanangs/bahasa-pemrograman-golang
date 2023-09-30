package main

import (
	"embed"
	"fmt"
	"os"
)

//go:embed embed.txt
var embedText string

//go:embed Screenshot.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {

	// string
	fmt.Println(embedText)

	// []byte
	err := os.WriteFile("logo_new.png", logo, 0666)
	if err != nil {
		panic(err)
	}

	// multiple files match
	dirEntry, _ := path.ReadDir("files")
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
