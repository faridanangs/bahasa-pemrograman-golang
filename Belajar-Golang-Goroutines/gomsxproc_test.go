package belajargolanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}

	// mengubah nilai goroutine
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total cpu:", totalCpu)

	// untuk mengubah nilai thread cukup tambahkan nilai yang di inginkan di dalam GOMAXPROCS()
	totalThread := runtime.GOMAXPROCS(-1)       // jika kita gunakan -1 di akam mengambil semua nilai
	totalThreadChange := runtime.GOMAXPROCS(20) // jika kita gunakan nlai > 0 maka dia akan menyetel niai threadnya sesuai dengan yang di masukan
	fmt.Println("Total thread:", totalThread)
	fmt.Println("Total thread:", totalThreadChange)

	// untuk mengubah nilai goroutine kita bisa gunakan for
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine", totalGoroutine)

	wg.Wait()
}
