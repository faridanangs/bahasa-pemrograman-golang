package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, Country, PostalCode string
}

type Customer struct {
	FirstName, LastName string
	Age                 int
	Hobbies             []string
	// kita memasukan data addres ke dalam slice supaya adress ini bisa di panggil sebanyak"nya
	// // kita gunakan slice supaya bisa di ubah menjadi array oleh json
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	// kita gunakan strct untuk memasukannya ke dalam json dan json mengubahnya menjadi object
	cust := Customer{
		FirstName: "Farid",
		LastName:  "Samudra",
		Age:       17,
	}

	bytes, _ := json.Marshal(cust)
	fmt.Println(string(bytes))
}
