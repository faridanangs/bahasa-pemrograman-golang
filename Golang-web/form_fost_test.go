package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(write http.ResponseWriter, request *http.Request) {
	// sebelum kita memanggil content di dlm body kita parse dulu menggunakan request.ParseForm()
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	// setelah di parse baru kita bisa panggil konten yang ada di dalam body menggunakan
	// request.PostForm.Get(key) key ini mengambil dare reader yang ada di dalam sini bodyForm := strings.NewReader("first_name=farid&last_name=anang")
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	fmt.Fprintf(write, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	// untuk bisa megirim nilI BODI ke header kita buat dulu reader menggunakan strings.NewReader("value")
	bodyForm := strings.NewReader("first_name=farid&last_name=anang")
	// kita masukan nilai bodyForm ke dakam body melalui httptest.NewRequest()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bodyForm)
	// supaya bisa menggunakan form post kita harus menambahkan ("Content-Type", "application/x-www-form-urlencoded") ke dalam header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)
	fmt.Println(string(body))

}
