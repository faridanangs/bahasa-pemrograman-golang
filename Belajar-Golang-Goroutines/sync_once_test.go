package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var count int = 0

func OnlyOne() {
	count++
}

func TestSyncOnce(t *testing.T) {
	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			once.Do(OnlyOne) // jika kita ingin sebuah function hanya di jalankan sekali saja kita bisa menggunakan sync.Once
		}()
		wg.Done()
	}

	wg.Wait()
	fmt.Println("Count", count)
}
