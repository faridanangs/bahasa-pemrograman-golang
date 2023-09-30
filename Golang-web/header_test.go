package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(write http.ResponseWriter, request *http.Request) {
	// untuk memangil header di request kita bisa gunakan request.Header.Get(key)
	headerType := request.Header.Get("content-type")
	fmt.Fprint(write, headerType)
}
func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	// untuk menambah header kita bisa menggunakakn Header.Add(key, value)
	// nilai dari header tidak peduli besar kecil jika kita memanggilnya
	request.Header.Add("Content-Type", "applicaion/json")
	recode := httptest.NewRecorder()

	RequestHeader(recode, request)

	response := recode.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
