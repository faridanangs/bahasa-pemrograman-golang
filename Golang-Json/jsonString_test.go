package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Farid",
		LastName:  "Samudra",
		Age:       17,
		Hobbies:   []string{"cooding", "running", "workout"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Farid","LastName":"Samudra","Age":17, "Hobbies":["cooding", "running", "workout"]}`

	customer := &Customer{}

	json.Unmarshal([]byte(jsonString), customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

func TestJsonArraComplex(t *testing.T) {
	customer := &Customer{
		FirstName: "Farid",
		LastName:  "Samudra",
		Addresses: []Address{
			{
				Street:     "Sesela 01 jakarta",
				Country:    "Indonesia",
				PostalCode: "382736",
			},
			{
				Street:     "dupa 01 jakarta",
				Country:    "Lombok",
				PostalCode: "343213",
			},
			{
				Street:     "kuku 01 jakarta",
				Country:    "Nesia",
				PostalCode: "7756732",
			},
		}}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArraComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Farid","LastName":"Samudra","Age":0,"Hobbies":null,"Addresses":[{"Street":"Sesela 01 jakarta","Country":"Indonesia","PostalCode":"382736"},{"Street":"dupa 01 jakarta","Country":"Lombok","PostalCode":"343213"},{"Street":"kuku 01 jakarta","Country":"Nesia","PostalCode":"7756732"}]}`
	customer := &Customer{}
	json.Unmarshal([]byte(jsonString), customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)

}

// kita juga bisa hanya memanggil stringnya saja tanpa ikut memanggil objectnya
func TestOnlyJsonArraComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Sesela 01 jakarta","Country":"Indonesia","PostalCode":"382736"},{"Street":"Sesela 01 jakarta","Country":"Indonesia","PostalCode":"382736"},{"Street":"Sesela 01 jakarta","Country":"Indonesia","PostalCode":"382736"}]`
	// kita gunakan type slice supaya addres bisa menerima lebih dari 1 data
	addresses := &[]Address{}
	json.Unmarshal([]byte(jsonString), addresses)

	fmt.Println(addresses)

}

func TestOnlyJsonArraComplex(t *testing.T) {
	addresses := &[]Address{
		{
			Street:     "Sesela 01 jakarta",
			Country:    "Indonesia",
			PostalCode: "382736",
		},
		{
			Street:     "dupa 01 jakarta",
			Country:    "Lombok",
			PostalCode: "343213",
		},
		{
			Street:     "kuku 01 jakarta",
			Country:    "Nesia",
			PostalCode: "7756732",
		},
	}

	bytes, _ := json.Marshal(addresses)

	fmt.Println(string(bytes))
}
