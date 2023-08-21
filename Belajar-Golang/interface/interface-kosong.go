package main

import "fmt"

// interface{} ini dapat menerima semua tipe data mulai dari int
// string dan boll dan juga mereturn semua jenis tipe data
// ini di sebut interface kosong

func upps(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	// kita tidak bisa menaruh tipe string/int/bool di dalam var inter
	// karna di function upps kita sudah mendeklarasikan interface{}
	// dia beisa menerima semua tipe data tanpa perlu di deklarasikan
	// terlwbih dahulu
	var inter interface{} = upps(1)
	fmt.Println(inter)
}
