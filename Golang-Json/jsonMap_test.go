package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// utuk menangkap nilai yang tak tentu di dalam json kita bisa menggunakan map[string]interfce{}

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Lolipop","image_url":"http://localhost/gambar.png"}`

	var product map[string]interface{}

	json.Unmarshal([]byte(jsonString), &product)
	fmt.Println(product)
	fmt.Println(product["id"])
	fmt.Println(product["name"])
	fmt.Println(product["image_url"])

}
func TestJsonMapEncode(t *testing.T) {
	// jika kita tak menaruh niai apapun yang di kembalikan struct kosong
	product := map[string]interface{}{}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
