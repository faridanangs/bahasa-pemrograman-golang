package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// func TestSyncMutex(t *testing.T) {
// 	// jika kita tidak menggunakan mutex maka hasil dari x tidak akan sama dengan
// 	// yang di harapkan karna dia akan menjalankan nilai yang sama pada waktu sang sama sehingga
// 	// nilai yang di teria berkurang
// 	// untuk mengatasi itu kita gunakan async.Mutex di mana kita me lock dan unlock datanya
// 	// kita masukan semua nilai ke dalam unlock() jika nilai sudah masuk ke dalam x
// 	// maka kita men unlocknya supaya nilai yang ada di dalam lock bisa di masukan ke dalam x
// 	//// ibarat harus mengantri giliran ok
// 	var mutex sync.Mutex
// 	x := 0
// 	for i := 1; i <= 1000; i++ {
// 		go func() {
// 			for j := 1; j <= 50; j++ {
// 				mutex.Lock()
// 				x = x + 1
// 				mutex.Unlock()
// 			}
// 		}()
// 	}
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("Counter ke ", x)
// }

// type BankAccount struct {
// 	RWmutex sync.RWMutex
// 	Balance int
// }

// func (account *BankAccount) AddBalance(count int) {
// 	account.RWmutex.Lock() // jika kita tidak menggunakan Mutex.Lock() maka dia akan terjadi cras di mana nilainya saling timpa
// 	account.Balance = account.Balance + count
// 	account.RWmutex.Unlock()
// }

// func (account *BankAccount) GetBalance() int {
// 	account.RWmutex.RLock() // kemudian kita gunakan rUnlock() untuk membaca nilai yang sudah di kirim ke dakam balance
// 	balance := account.Balance
// 	account.RWmutex.RUnlock()
// 	return balance
// }

// func TestRwMutex(t *testing.T) {
// 	var wg sync.WaitGroup
// 	account := BankAccount{}

// 	for i := 0; i <= 1000; i++ {
// 		wg.Add(1)
// 		go func(acc *BankAccount) {
// 			defer wg.Done()
// 			for j := 0; j <= 100; j++ {
// 				acc.AddBalance(1)
// 				fmt.Println(acc.GetBalance())
// 			}
// 		}(&account)
// 	}

// 	wg.Wait()
// 	fmt.Println("total balance", account.GetBalance())
// }

// Deadlock

// type UserBalance struct {
// 	sync.Mutex
// 	Name    string
// 	Balance int
// }

// func (user *UserBalance) Lock() {
// 	user.Mutex.Lock()
// }
// func (user *UserBalance) Unlock() {
// 	user.Mutex.Unlock()
// }

// func (user *UserBalance) Change(amount int) {
// 	user.Balance = user.Balance + amount
// }

// func Transfer(user1 *UserBalance, user2 *UserBalance, amount int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	// Mengunci pengguna dengan nilai lebih rendah terlebih dahulu
// 	if user1.Name < user2.Name {
// 		user1.Lock()
// 		user2.Lock()
// 	} else {
// 		user2.Lock()
// 		user1.Lock()
// 	}

// 	fmt.Println("Lock user 1", user1.Name)
// 	user1.Change(-amount)

// 	time.Sleep(3 * time.Second)

// 	fmt.Println("Lock user 2", user2.Name)
// 	user2.Change(amount)
// 	time.Sleep(3 * time.Second)

// 	user1.Unlock()
// 	user2.Unlock()
// }

// func TestDeadlock(t *testing.T) {
// 	user1 := UserBalance{
// 		Name:    "Farid",
// 		Balance: 1000000,
// 	}
// 	user2 := UserBalance{
// 		Name:    "Anans",
// 		Balance: 1000000,
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go Transfer(&user1, &user2, 100000, &wg)
// 	go Transfer(&user2, &user1, 200000, &wg)

// 	wg.Wait()

// 	fmt.Println("User ", user1.Name, "Balance", user1.Balance)
// 	fmt.Println("User ", user2.Name, "Balance", user2.Balance)
// }

func RunAyncRonus(grup *sync.WaitGroup) {
	defer grup.Done()
	fmt.Println("Hello world")
	time.Sleep(4 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go RunAyncRonus(&wg)
	}

	wg.Wait()
	fmt.Println("selesai")

}
