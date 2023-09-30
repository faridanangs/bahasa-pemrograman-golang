package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}
func TestAfterTimer(t *testing.T) {
	timer := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer
	fmt.Println(time)
}

func TestFuncTimer(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	// kita bisa membuat sebah function supaya di jalankan pada waktu tertentu menggunakan AfterFunc()
	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println("data sudah masuk ke dalam rekening anda", time.Now())
	})

	fmt.Println("Data sedang di kirim ke rekening anda", time.Now())
	wg.Wait()
}

func TestNewTicker(t *testing.T) {
	timer := time.NewTicker(1 * time.Second) // kita gunakan NewTicker untuk membuat waktu secara terus menerus

	go func() {
		time.Sleep(5 * time.Second)
		timer.Stop() // jika kita tidak menggunakan stop maka timer ini akan terus berjalaln tanpa henti
	}()

	for data := range timer.C {
		fmt.Println(data)
	}
}

func TestTick(t *testing.T) {
	stop := make(chan bool)

	timeTicker := time.Tick(1 * time.Second)
	go func() {
		for {
			select {
			case <-timeTicker:
				fmt.Println(time.Now().Clock())
			case <-stop:
				return // Goroutine berhenti saat konteks dibatalkan
			}
		}
	}()
	// Biarkan goroutine berjalan selama beberapa waktu sebelum dihentikan
	time.Sleep(5 * time.Second)
	stop <- true
}
