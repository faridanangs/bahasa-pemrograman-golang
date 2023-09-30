package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(write http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {

		// untuk bisa mengirim status code ke dalam header kita bisa menggunakan write.WriteHeader(int/statusCode)
		write.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(write, "Name is empty")
	} else {
		write.WriteHeader(http.StatusOK)
		fmt.Fprintf(write, "hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()

	ResponseCode(recode, request)

	response := recode.Result()
	body, _ := io.ReadAll(response.Body)

	// untuk bisa melihat status code berupa int kita bisa menggunkan .StatusCode
	fmt.Println(response.StatusCode)
	// untuk bisa melihat status code berupa string kita bisa menggunkan .Status
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080?name=farid", nil)
	recode := httptest.NewRecorder()

	ResponseCode(recode, request)

	response := recode.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
