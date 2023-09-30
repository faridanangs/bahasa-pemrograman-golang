package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// func TestWithValue(t *testing.T) {
// 	contextA := context.Background()
// 	// UNTUK memasukan nilai ke dalam context kita bisa gunakan WithValue(partn, key, value)
// 	contextB := context.WithValue(contextA, "b", "Bebe")
// 	contextC := context.WithValue(contextA, "c", "Cebe")
// 	contextD := context.WithValue(contextB, "d", "Debe")
// 	contextE := context.WithValue(contextB, "e", "Eebe")
// 	contextF := context.WithValue(contextC, "f", "Febe")
// 	// jika kita ingin menambah/mengubah nilai yang ada di dalam context sebenarnya kita tidak mengubahnya
// 	// melainkan kita membuat child baru yang di mana dia akan bersusun
// 	// contohnya seperti ini
// 	// context.Background.WithValue(type string, val Bebe).WithValue(type string, val Debe)
// 	// dimana contextD anak dari contextB dan contextB anak dari contextA
// 	fmt.Println(contextA)
// 	fmt.Println(contextB)
// 	fmt.Println(contextC)
// 	fmt.Println(contextD)
// 	fmt.Println(contextE)
// 	fmt.Println(contextF)
// 	// cara memanggil nilai yang ada di dalam context kita bisa menggunakan Valu(key)
// 	// jika nilai sesuai dgn yang ada di dalamnya maka dia akan memanggilnya jika tidak ada
// 	// dia akan naik ke parntnya untuk memanggilnya jika tidak ada dia naik lagi sampai ke parent indux terakhir
// 	// jika tidak ada dia akan mengembalikan nil
// 	fmt.Println(contextF.Value("b"))
// 	fmt.Println(contextF.Value("c"))
// 	fmt.Println(contextF.Value("a"))
// 	fmt.Println(contextA.Value("d"))
// }

// WithCancel
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("data ke", n)
		if n == 10 {
			break
		}
	}

	cancel() // cancel ini memberi sinyak ke contextWithCencel untuk memberhentikan contex kemudian akan di respone oleh ctx dan menjalankan done() kemudian berhenti dengan
	// cara me return kosong
	time.Sleep(1 * time.Second)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel() // walaupun kita sudah memanggiil cancel menggunakan WithTimeout secata otomatis menggunakan waktu kita harus gunakan cancel ini untuk jaga jaga jika terjadi kesalahan

	destination := CreateCounter(ctx)

	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("data ke", n)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())

	parent := context.Background()

	// sama dengan withTimeout, withDeadline hanya beda di waktunya saja seperti waktu di deadline menggunakan waktu yang akan datang seperti nanti jam 12 siang kah dll
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(time.Hour))
	defer cancel() // walaupun kita sudah memanggiil cancel menggunakan WithTimeout secata otomatis menggunakan waktu kita harus gunakan cancel ini untuk jaga jaga jika terjadi kesalahan

	destination := CreateCounter(ctx)

	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("data ke", n)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
}
