package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	nameFile := r.URL.Query().Get("file")
	if nameFile == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "status bad Request")
		return
	}

	// penjelasan tentang Content-Disposition
	{ /*
			"attachment; filename=""+nameFile+""":

			Dalam contoh ini, header "Content-Disposition" diatur dengan opsi "attachment", yang mengindikasikan bahwa file yang dikirimkan oleh server adalah file yang akan diunduh oleh pengguna.
			Opsi "filename" digunakan untuk menentukan nama file yang akan diberikan pada file yang diunduh. Nama file ini diambil dari variabel nameFile yang Anda dapatkan dari query parameter URL atau sesuai dengan kebutuhan Anda.
			Ketika pengguna mengklik tautan atau mencoba membuka file, browser biasanya akan memunculkan dialog "Simpan File" dan menggunakan nama yang telah ditentukan sebagai nama file default.
			"attachment":

			Dalam kasus ini, hanya opsi "attachment" yang diatur dalam header "Content-Disposition".
			Ini mengindikasikan bahwa file yang dikirimkan oleh server adalah file yang akan diunduh, tetapi tanpa memberikan nama file default. Oleh karena itu, browser akan menampilkan dialog "Simpan File" dengan nama file default yang harus diatur oleh pengguna.
			"inline":

			Header "Content-Disposition" diatur dengan opsi "inline".
			Ini mengindikasikan bahwa file yang dikirimkan oleh server seharusnya tidak diunduh secara otomatis, melainkan ditampilkan langsung oleh browser atau program yang sesuai. Ini cocok untuk menampilkan gambar atau file dalam format yang dapat ditampilkan langsung dalam browser, seperti PDF atau file teks.
			Ketika opsi "inline" digunakan, file akan dibuka atau ditampilkan dalam jendela browser tanpa mengharuskan pengguna untuk mengunduhnya terlebih dahulu.
		*/
	}

	// kita menggunakan Content-Disposition dan params attachment untuk membuat user mendownload file img yang dia ketik di url dengan nama file
	// w.Header().Add("Content-Disposition", "attachment; filename=\""+nameFile+"\"")

	// jika kita menggunakan attachment tanpa filename
	// w.Header().Add("Content-Disposition", "attachment")
	w.Header().Add("Content-Disposition", "inline")
	http.ServeFile(w, r, "./resorces/"+nameFile)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	server.ListenAndServe()
}
