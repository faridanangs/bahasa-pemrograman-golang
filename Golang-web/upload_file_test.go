package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func FormFile(w http.ResponseWriter, r *http.Request) {
	htmlFormFile := template.Must(template.ParseFiles("./templates/form_post.gohtml"))
	htmlFormFile.ExecuteTemplate(w, "form_post.gohtml", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// di FormFile kita memasukan nama dari name yang ada di file html form_post.gohtml
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// disini kita membuat file baru menggunakan os.Create di dalam folder resorces yang mengambil nama file dari fileHeader.Filename di mana name dari FileName  ini akan memangambil
	// dari file yang kita masukan i html
	fileDest, err := os.Create("./resorces/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// disini kita memasukan nilai file ke dalam fileDesc menggunakan io.Copy
	_, err = io.Copy(fileDest, file)
	if err != nil {
		panic(err)
	}

	// kita mengambil nilai name yang yang sudah di kirim di html
	name := r.PostFormValue("name")
	htmlFormFile := template.Must(template.ParseFiles("./templates/upload_success.gohtml"))
	htmlFormFile.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestFormFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", FormFile)
	mux.HandleFunc("/upload", UploadFile)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("/resorces"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

//go:embed resorces/alquran.jpg
var testUploadFile []byte

// testign upload file
func TestUploadFile(t *testing.T) {

	// untuk membuah unit test kita buat dulu body menggunakan new(bytes.Buffer)
	body := new(bytes.Buffer)

	// kemudian kita memasukan body ke dalam multipart.NewWriter
	writer := multipart.NewWriter(body)

	// kemudian di sini kita masuk ke name="name" di dalam form html menggunakan writer.CreateFormFile dan memasukan namenya "aNANGS"
	writer.WriteField("name", "aNANGS")
	// dan di sini kita masuk ke name="file" dan memasukan nama filenya "FileLoliTest.jpg"
	file, _ := writer.CreateFormFile("file", "FileLoliTest.jpg")

	// setelah itu kita masukan file yang sudah kita embed ke dalam file.Write(namaFilenya)
	file.Write(testUploadFile)
	// setelah itu kita harus clos supaya tidak EOF
	writer.Close()

	// supaya ini bisa berjalan kita harus menggunakan MethodPost dan memasukan nilai dari body ke dalam NewRequest
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/upload", body)
	// setelah itu kita harus menemtukan content typenya bisa menggunakan "multipart/form-data" bisa juga seperti ini writer.FormDataContentType()
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recode := httptest.NewRecorder()
	UploadFile(recode, request)

	bodyReq, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(bodyReq))
}
