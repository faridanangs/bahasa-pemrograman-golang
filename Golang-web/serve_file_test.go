package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// servefile tidak bisa menggunakan embed namun kita bisa menggunakan embed jika untuk melakukan load file
// kita hanya butuuh package fmt dan responseWriter
func ServeFile(write http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(write, request, "./resorces/ok.html")
	} else {
		http.ServeFile(write, request, "./resorces/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	server.ListenAndServe()
}

// serveFile embed

// kita ubah nilai yang ada di dalam file menggunaan embed menjadi string
//
//go:embed resorces/ok.html
var embedOk string

//go:embed resorces/notfound.html
var embedNotfound string

func ServeFileEmbed(write http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		// kita masukan nilai dari embed berupa string ke dalam Fprint()
		fmt.Fprint(write, embedOk)
	} else {
		fmt.Fprint(write, embedNotfound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		// kita panggil func ServeFileEmbed di dalam Handler
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	server.ListenAndServe()
}
