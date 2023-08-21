package main

import "fmt"

// differ
func userLogin() {
	fmt.Println("Function defer sudah di jalankna")
}

func funcJalankan(val int) {
	defer userLogin()
	fmt.Println("Function running")
	result := 10 / val
	fmt.Println(result)
}

func endApp() {
	// kita gunakan recover untuk menampung nilai dari panic
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error pada ", message)
	}
	fmt.Println("Aplikasi di jalankan")

}

func runApp(err bool) {
	// deffer ini di gunakan untuk memanggil function dan harus di taruh di paling
	// atas supaya aman
	defer endApp()
	if err {
		// panic ini di gunakan untuk memberhentikan kode yang sedang berjalan
		// secara tiba" tidak peduli dengan yng ada di bawahnya
		// jika terjadi panic dia akan men throw ke atas sehingga kita
		// bisa menggunakan/memanggil recover di atas panic seperti di func defer
		panic("Aplikasi ErrOR sayang")
	}
	fmt.Println("aplikasi selesai")
}

func main() {
	// funcJalankan(0)
	runApp(false)

	// jika kita tidak menaruh recover() maka fmt di bawah ini tidak akan berjalan
	// karna ketika panic di jalankan kodenya langsung berhenti dan tidak
	// peduli dengan yang lain
	fmt.Println("anangs")
}
