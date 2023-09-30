package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// supaya nama dari struct ketika di encode ke json berubah kita bisa menggunakan tag json seperti di bawah ini
type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Lolipop",
		ImageUrl: "http://localhost/gambar.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}
func TestJsonTagDecode(t *testing.T) {
	// jika kita ubah nama dari jsonnya golang tidak memperdulikannya asal sesuai dengan jsonTAg dan namanya
	jsonString := `{"ID":"P001","NAME":"Apple Lolipop","image_url":"http://localhost/gambar.png"}`

	product := &Product{}

	json.Unmarshal([]byte(jsonString), product)
	fmt.Println(product)

}
