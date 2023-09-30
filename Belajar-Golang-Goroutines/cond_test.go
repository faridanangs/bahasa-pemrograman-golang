package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var loker = sync.Mutex{}
var cond = sync.NewCond(&loker)
var wg sync.WaitGroup

func WaitCond(value int) {
	defer wg.Done()

	cond.L.Lock()
	cond.Wait() // kita gunakan wait untuk mencegah kode di jalankan sampai mendapat perintah untuk di jalankan melalui Signal dan Broadcast
	fmt.Println("Done", value)

	cond.L.Unlock()
}
func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WaitCond(i)
	}

	go func() {
		for j := 0; j < 10; j++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // Signal() ini memberi tau cond bahwa jalankan goroutine secara satu satu melalui wait
		}
	}()
	// go func() {
	// 	time.Sleep(1 * time.Second) // jika kita tdk menaruh slep maka kodenya akan terjadi tras
	// 	cond.Broadcast() // jika kita gunakan Broadcast() maka semua gorutine akan di jalankan tanpa menunggu wait
	// }()

	wg.Wait()
}
