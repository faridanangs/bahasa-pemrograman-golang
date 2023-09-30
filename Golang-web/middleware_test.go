package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	// karna disini kita gunakan Type http.Handler maka nama funct yang melakukan
	// kontrak harus ServerHTTP(write, req)
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(write http.ResponseWriter, req *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(write, req)
	fmt.Println("After execute handler")

	{ /*
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println("Terjadi error")
					write.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(write, "Error %s", err)
				}
			}()
			fmt.Println("After execute error")
		*/
	}

	// jika kode ini kita taruh di atas maka kode error tidak bisa di jalankan maka
	//  dari itu kita buat func baru husus untuk menangkap error
	////// middleware.Handler.ServeHTTP(write, req)
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(write http.ResponseWriter, req *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			write.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(write, "Error %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(write, req)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handler Execute")
		fmt.Fprint(w, "Hello Middleware")
	})

	// untuk menangkap error supaya program kita tidak berhenti kita gunakan panic dan recover untuk menangkap error tersebut
	// dan membuat struct husus untuk menangani error
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Ups")
	})

	// kemudian logMiddleware akan menjalankan mux
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// ketika error handler di jalankan dan tidak terjadi error maka dia akan merequest ke logmiddleware
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	// ketika terjadi request maka server di jalankan kemudian server menjalankan errorHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}
	server.ListenAndServe()
}
