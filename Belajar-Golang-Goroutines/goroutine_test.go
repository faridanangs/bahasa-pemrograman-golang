package belajargolanggoroutines

import (
	"fmt"
	"testing"
)

// func HelloWorld() {
// 	fmt.Println("Hello World")
// }

// func TestGoRountine(t *testing.T) {
// 	// go routine ini berjalan seperti async di js di mana dia tidak akan
// 	// menunggu yang lain selesai tetapi dia akan langsung menjalankan yang lain dulu
// 	go HelloWorld()
// 	fmt.Println("Upsss")
// 	time.Sleep(1 * time.Second)
// }

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGorountine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
}

// 45,29s tanpa gorountine
// 10.000 = 0,30s
