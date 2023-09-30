package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamerEncode(t *testing.T) {
	// nama file yang akan kita buat menjadi json
	reader, _ := os.Create("customer_2.json")
	// kemudian kita encoder dulu filel jsonnya
	streamEncode := json.NewEncoder(reader)

	customer := &Customer{
		FirstName: "farid",
		LastName:  "anangs",
		Addresses: []Address{
			{
				Street:     "sesela 02 guru",
				Country:    "indonesia",
				PostalCode: "830730231",
			},
			{
				Street:     "kbon 02 kisam",
				Country:    "indonesia",
				PostalCode: "83313231",
			},
			{
				Street:     "bakau 043 koki",
				Country:    "indonesia",
				PostalCode: "731738",
			},
			{
				Street:     "gurut 02 mulkis",
				Country:    "indongogi",
				PostalCode: "37371030",
			},
		},
	}

	// kemudian kita masukan data customer ke dalam json menggunakan streamEncode.Encode(customer)
	streamEncode.Encode(customer)
	fmt.Print(customer)

}

func TestAmbilDataJsonYangSudahDiKirim(t *testing.T) {
	reader, _ := os.Open("customer_2.json")

	decod := json.NewDecoder(reader)

	customer := &Customer{}
	decod.Decode(customer)
	fmt.Println(customer)
}
