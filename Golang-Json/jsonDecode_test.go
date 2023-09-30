package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	jsonString := `{"FirstName": "Farid","LastName":  "Samudra","Age":17}`

	customer := &Customer{}

	// untuk mengubah data json menjadi struct kita bisa menggunakan json.Unmarshal([]byte(), interface{})
	// interface{} ini tempat data hasil dari decode di simpan
	json.Unmarshal([]byte(jsonString), customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)

}
