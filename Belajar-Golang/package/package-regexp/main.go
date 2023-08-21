package main

import (
	"fmt"
	"regexp"
)

func main() {
	// kita gunakan regexp.MustCompile untuk menginisialisasikan hurup yang akan kita coba
	var regex = regexp.MustCompile(`f([a-z])[a-z][a-z]d`)
	re := regexp.MustCompile(`f.`)

	// kita gunakan regex.MatchString untuk mencoba kalimat yang ingin kita cari dengan mengembaliakn
	// boolean jika kalimatnya ada di dalam regex
	fmt.Println(regex.MatchString("farid"))
	fmt.Println(regex.MatchString("fukid"))
	fmt.Println(regex.MatchString("fusut"))

	// kita gunakan FindAllString untuk mencari/mencocokan kalimat yang ada di dalam
	// regex jika cocok maka akan di masukan ke dalam slice jika tidak dia tidak
	// di masukan ke dalam slice
	content := []byte("London")
	fmt.Println(regex.FindAllString(`farid fakid fukud fukud fakir`, -1))
	fmt.Println(re.FindAllIndex(content, -1))

}
