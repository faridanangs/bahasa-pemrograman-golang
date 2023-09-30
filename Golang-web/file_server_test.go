package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	// kita gunakan http.Dir untuk mengambil nama folder di aplikasi kita
	directory := http.Dir("./resorces")

	// kemudian kita gunakan http.FileServer() supaya bisa di tampilkan di server
	fileServe := http.FileServer(directory)

	// kita buat variable mux
	mux := http.NewServeMux()
	// setelah itu kita masukan kedalam mux supaya bisa di panggil di handle
	// kemudian kita gunakan http.StripPrefix() supaya bisa di panggil di url dengan mudah tanpa memanggil name foldernya atau
	//  Penggunaan http.StripPrefix digunakan untuk menghapus awalan "/static/" dari URL permintaan sehingga server dapat mencari file yang sesuai dalam direktori "resorces".
	mux.Handle("/statis/", http.StripPrefix("/statis/", fileServe))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// jika kita menggunakan embed maka folder di dalam akan di kompile dan di masukan ke dalam variabel
// dan ini tentunya lebih bagus dan cepat tanpa perlu mengikutkan foldernya jika kita build nanti

//go:embed resorces

// ini aadalah deklarasi variable global yang bertipe embed.FS
// variablel ini akan di gunakan untuk mengakses directory resorces yang tellah di embed
var resorc embed.FS

func TestFileServerEmbed(t *testing.T) {

	// ini adallah baris kode yang mengambil sub directory dari resorc yang ada di dalam resorc
	direc, _ := fs.Sub(resorc, "resorces")

	// kita gunakan http.FS untuk melayani kode kode yang sudh di embed
	// jika kita tidak menggunakan ini http.FS maka direc ini tidak bisa di baca oleh FileServer karna ini adalah hasil dari embed
	fileServer := http.FileServer(http.FS(direc))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
