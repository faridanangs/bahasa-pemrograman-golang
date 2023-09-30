package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(recode http.ResponseWriter, r *http.Request) {

	// kita ambil nilai name dalam url menggunakan r.URL.Query().Get("name")
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(recode, "Hello")
	} else {
		fmt.Fprintf(recode, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	// kita membuath request test supaya bisa di gunakan di func SayHello dan mengirim request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=anangs", nil)
	// kita menggunakan recorder untuk menanggkap hasil dari func SayHello dan di simpan di recorder
	recode := httptest.NewRecorder()

	SayHello(recode, request)

	// kita abil nilai dari recode yang masih berupa response
	response := recode.Result()

	// kita parse nilai dari respone yang masih berupa response menjadi []byte
	body, _ := io.ReadAll(response.Body)

	// kita ubah nilai dari body yang beupa byte menjadi string
	fmt.Println(string(body))
}

func MultipleParams(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
}

func TestMultipleParams(t *testing.T) {
	// jika kita ingin membuat TestMultipleParams cukup tarus & di dalam urlnya
	request := httptest.NewRequest(http.MethodGet, "http://localhost/sapa?first_name=farid&last_name=anang", nil)
	recode := httptest.NewRecorder()
	MultipleParams(recode, request)

	response := recode.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParamsValues(w http.ResponseWriter, r *http.Request) {
	// jika kita menggunakan r.URL.Query().GET("name")
	// maka kita hanya akan mendapatkan nilai name yang paling depan

	// jika kita menggunakan r.URL.Query()["name"]
	// maka kita akan mendapatkan semua nilai name
	// karna dia merupakan map[string][]string

	query := r.URL.Query()["name"]
	fmt.Fprintf(w, strings.Join(query, " "))
}

func TestMultipleParamsValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/sapa?name=farid&name=anang&name=samudra", nil)
	recode := httptest.NewRecorder()
	MultipleParamsValues(recode, request)

	response := recode.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
