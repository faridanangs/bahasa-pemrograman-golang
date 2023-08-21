package main

import (
	"flag"
	"fmt"
)

func main() {
	// jika kita tidak masukan nilai maka dia akan di set dengan nilai default
	// yang di beri
	// cara manggil namanya kita gunakan -namanya di terminal
	// // flag.String("namanya apa", "default-nilainya apa", "textnya apa")

	// kita coba dengan Variable
	// supaya bisa menggunakan variable kita gunakan flag.StringVar(var, "nama", "default", "text")
	var dataBuku string
	const (
		defaultText = "ikn"
		text        = "Masukan nama bukunya"
	)
	flag.StringVar(&dataBuku, "buku", defaultText, text)

	host := flag.String("host", "localhost", "Put your localhost to data base")
	username := flag.String("user", "root", "Create your name user to data base")
	password := flag.String("password", "root", "Create your password to data base")

	// untuk bisa menampilkan nilai hasil perubahan kita gunakan Parse()
	flag.Parse()

	// Supaya println dapat menampilkan hasilnya kita gunakan pointer * di awal nama variable
	fmt.Println("Localhost: ", *host)
	fmt.Println("Username: ", *username)
	fmt.Println("Password: ", *password)
	fmt.Println("Buku: ", dataBuku)

}
