package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamerDecoder(t *testing.T) {
	// kita ambil file jsonnya
	reader, _ := os.Open("customer.json")
	// kemudian kita masukan data json ke dalam NewDecoder()
	decod := json.NewDecoder(reader)

	// harus berupa pointer
	customer := &Customer{}

	// setalh di masukan ke dalam NewRecorder kemudian kita Decode sama seperti Marshal() untuk mendapatkan data yang ada di dalamnya
	decod.Decode(customer)

	fmt.Println(customer)
}
