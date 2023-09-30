package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	pool := sync.Pool{
		New: func() any {
			return ""
		},
	}
	// kita gunakan put untuk memasukan data ke dalam pool
	pool.Put("Farid")
	pool.Put("Anang")
	pool.Put("Samudra")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get() // kita gunakan get untuk mengambil data di dalam pool
			fmt.Println(data)
			pool.Put(data) // setelah kita ambil datanya kita masukan kembali data yang sudah di ambil ke dalam pool
		}()
	}
	wg.Wait()
	fmt.Println("selesai", pool.Get(), pool.Get(), pool.Get())
}
