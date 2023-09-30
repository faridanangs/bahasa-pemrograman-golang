package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// jika kita gunakan http.HandlerFunc maka ini url pointernya hanya satu tidak bisa lebih
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "<h1>Hello dunia</h1>")
	}

	server := http.Server{
		Addr: "localhost:8080",

		// kita gunakan handle untuk menampulkan konten ke client
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestMuxServer(t *testing.T) {
	// kita gunakan http.NewServeMux() untuk menampung semua HandleFunc yang di berikan
	// jika kita gunakan http.NewServeMux() maka ini url pointernya bisa lebih dari 1
	// dan kita gunakan HandleFunc("url", "kontent")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "<h1>Hello dunia</h1>")
	})
	mux.HandleFunc("/sapa", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "<h1>Hello User</h1>")
	})

	// jika kita menaruh / di akhir url maka apapaun nilai di belakang / akan di masukan ke dalam parentnya
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "<h1>Hello Images</h1>")
	})
	mux.HandleFunc("/images/thumbnail/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "<h1>Hello Thumbnail</h1>")
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
