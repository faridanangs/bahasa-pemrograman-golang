package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	// untuk melakukan encode kode golang ke json kita cukup gunakan json.Marshal(interface{})
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonEncode(t *testing.T) {
	logJson("anangs")
	logJson(1)
	logJson(true)
	logJson([]string{"far", "anang", "samud"})
}
