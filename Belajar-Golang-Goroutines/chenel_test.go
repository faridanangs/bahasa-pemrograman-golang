package belajargolanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// func TestChanel(t *testing.T) {

// 	// chanel berjalan secara syincronus di mana dia akan menerima satu data
// 	// dan akan mengirim satu data jika data yang ada di dalam chanel tidak di ambil
// 	// maka chalen tdk akan menerima data baru atau block
// 	// cara membuat chanel kita gunakan make(chan type-data)
// 	chanel := make(chan string)
// 	// cara memasukan data ke cenel cukup menggunakan tanda <-
// 	// chanel <- "data"
// 	// jangan lupa untuk menclose chanel supaya datanya tidak flak
// 	defer close(chanel)
// 	// kita gnakan go routine dan func anonim untuk mengirim data ke chanel
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		// cara menerima dan mengirim chanel cukup gunakan <- chanel
// 		chanel <- "Farid anang samudra"
// 		fmt.Println("Selesai mengirim data ke dalam chanel")
// 	}()
// 	// cara menerima dan mengirim chanel cukup gunakan <- chanel
// 	data := <-chanel
// 	fmt.Println(data)
// 	time.Sleep(3 * time.Second)
// }

// chanel sebagai parameter
// func chanalFunction(dataChan chan string) {
// 	time.Sleep(3 * time.Second)
// 	dataChan <- "Farid anang samudra"
// }

// func TestChanelParameter(t *testing.T) {
// 	chanel := make(chan string)
// 	defer close(chanel)
// 	go chanalFunction(chanel)
// 	data := <-chanel
// 	fmt.Println(data)
// }

// func TestInOut(t *testing.T) {
// 	channel := make(chan string)
// 	defer close(channel)
// 	go OnlyOut(channel)
// 	go OnlyIn(channel)
// 	time.Sleep(5 * time.Second)
// }

// // untuk hanya memesukkan data ke dalam channel kita gunakan tnda param chan <- string
// func OnlyIn(channel chan<- string) {
// 	time.Sleep(3 * time.Second)
// 	channel <- "farid anang"
// 	channel <- "farid anang samudra"
// }

// // untuk hanya menerima data ke dalam channel kita gunakan tnda param <- chan string
// func OnlyOut(channel <-chan string) {
// 	data := <-channel
// 	data2 := <-channel
// 	fmt.Println(data)
// 	fmt.Println(data2)
// }

// buffer
// func TestBuffer(t *testing.T) {
// 	// kita gunakan buffer untuk menampung data yang di kirim walaupun
// 	// tidak di panggil channel tidak bloking karna kita sudah menampung
// 	// datanya ke dalam buffer.
// 	// cara memanggil buffer cukup taruh angka int di terakhir example
// 	//// make(chan string, buffer)
// 	channel := make(chan string, 3)
// 	defer close(channel)
// 	go func() {
// 		channel <- "farid"
// 		channel <- "anang"
// 		channel <- "samudra"
// 	}()
// 	go func() {
// 		fmt.Println(<-channel)
// 		fmt.Println(<-channel)
// 		fmt.Println(<-channel)
// 	}()
// 	fmt.Println("selesai")
// 	time.Sleep(5 * time.Second)
// }

// func TestRangeChanel(t *testing.T) {
// 	channel := make(chan string)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			channel <- "channel ke " + strconv.FormatInt(int64(i), 10)
// 		}
// 		// jika kita tidak menaruh close dia akan deadloock / error
// 		close(channel)
// 	}()
// 	// kita tidak perlu menggunakan <- untuk memanggil data yang ada di dalam chanel
// 	for data := range channel {
// 		fmt.Println(data)
// 	}
// 	fmt.Println("selseai")
// }

func Givemeresponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Farid anang samudra"
}

func TestChannelSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go Givemeresponse(channel1)
	go Givemeresponse(channel2)

	// jika kita hanya menggunakan select tanpa for maka yang akan di jalankan adalah yang
	// paling cepat saja atau yang pertama
	// jila kita mneggunakan for maka semua data di dalam channel akan di jalankan namun kita
	// harus membuat kondisi untuk mem breaj for tersebut supaya tidak deadloock
	count := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel ke 1", data)
			count++
		case data := <-channel2:
			fmt.Println("data dari channel ke 2", data)
			count++

		// jika data belum memiliki nilai makan akan terjadi error untuk mencegah terjadi error tersebut
		// kita gunakan default untuk di jalankan sampai data memiliki nilai
		default:
			fmt.Println("Menunggu data")
		}
		if count == 2 {
			break
		}
	}
}
