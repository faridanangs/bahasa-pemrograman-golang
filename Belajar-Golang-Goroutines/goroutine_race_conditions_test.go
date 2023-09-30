package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	// jika kita tidak menggunakan mutex maka hasil dari x tidak akan sama dengan
	// yang di harapkan karna dia akan menjalankan nilai yang sama pada waktu sang sama sehingga
	// nilai yang di teria berkurang
	// untuk mengatasi itu kita gunakan async.Mutex di mana kita me lock dan unlock datanya
	// kita masukan semua nilai ke dalam unlock() jika nilai sudah masuk ke dalam x
	// maka kita men unlocknya supaya nilai yang ada di dalam lock bisa di masukan ke dalam x
	//// ibarat harus mengantri giliran ok
	var mutex sync.Mutex
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 50; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(4 * time.Second)
	fmt.Println("Counter ke ", x)
}
