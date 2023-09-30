package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SetCookie(write http.ResponseWriter, request *http.Request) {
	// untuk membuat cookie kita harus mengunakan new(http.Cookie)
	cookie := new(http.Cookie)
	// kwmudian memasukan nama ke dalam cookie use cookie.Name
	cookie.Name = "X-Nangs"
	// kwmudian memasukan value ke dalam cookie use cookie.Value
	cookie.Value = request.URL.Query().Get("name")
	// kemudian untuk memasukan nilai dari cookie.name dan value kita bisa menggunakan http.SetCookie(write, cookie)
	http.SetCookie(write, cookie)
	fmt.Fprint(write, "succes create cookie")
}

func GetCookie(write http.ResponseWriter, request *http.Request) {
	// untuk mendapatkan nilai dari cookie kita bisa menggunakan request.Cookie("name")
	cookie, err := request.Cookie("X-Nangs")
	if err != nil {
		fmt.Fprint(write, "No cookie")
	} else {
		// untuk mendapatkan value dari cookie kita bisa mengguanakn cookie.Value
		values := strings.Split(cookie.Value, ",")

		// Iterasi dan cetak semua nilai
		for _, value := range values {
			fmt.Fprintf(write, "Value: %s\n", value)
		}
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-cookie", func(write http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("X-2Anangs")
		if err != nil {
			fmt.Fprint(write, "No cookie")
		} else {
			name := cookie.Value
			fmt.Fprintf(write, "Hello %s", name)
		}
	})
	mux.HandleFunc("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		cookie := new(http.Cookie)
		cookie.Name = "X-Nangs"
		cookie.Value = r.URL.Query().Get("name")
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		fmt.Fprint(w, "success create cookie")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/?name=farid", nil)
	recode := httptest.NewRecorder()

	SetCookie(recode, request)
	// untuk mendapatkan nilai dari cooki yang berupa slice kita bisa menggunakan recode.Result().Cookies()
	response := recode.Result().Cookies()
	for _, data := range response {
		fmt.Printf("Cookie %s %s \n", data.Name, data.Value)
	}

}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	// supaya ini bisa di panggil di get kita harus menamakan nama yang sama dengan nama cookie yang kita miliki
	cookie.Name = "X-Nangs"
	cookie.Value = "samudra"
	request.AddCookie(cookie)

	recode := httptest.NewRecorder()
	GetCookie(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))

}
