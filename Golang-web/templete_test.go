package golangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// template string
func HtmlSimple(write http.ResponseWriter, request *http.Request) {
	// kita buat dulu html template nya kemudian kita gunakan {{.}} untuk menangkap nilai
	// dinamis yang di kirim jika kita tidak menggunakan . di dalam {{}} maka dia akan error
	htmlText := "<html><body>{{}}</body></html>"
	// // setelah itu kita buat templatenya dan memparse htmltextnya
	// t, err := template.New("SIMPLE").Parse(htmlText)
	// if err != nil {
	// 	panic(err)
	// }
	// ini cara yang lebih simple kita tidak perlu menangkap errornya
	t := template.Must(template.New("SIMPLE").Parse(htmlText))
	// setelah itu kita gunakan ExecuteTemplate untuk memasukan nilia ke dalam name SIMPLE dan memasukan text yang
	// ingin kita tulis ke dalam htmlText jika nama dari templateNew beda dengan nama yang ada di ExecuteTemplate maka
	// dia akan error
	t.ExecuteTemplate(write, "SIMPLE", "INI html template simple")
}

func TestHtmlSimple(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()

	HtmlSimple(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}

// template file unutk bisa melakukan ini kita harus membuat file html dengan akhiran .gohtml
func SimpleHtmlFile(write http.ResponseWriter, request *http.Request) {
	// di sini hanya beda kita menggunakan file html dan memparsenya menggunakan ParseFiles dan tempat filenya di simpan
	htmlFile := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	htmlFile.ExecuteTemplate(write, "simple.gohtml", "ini html file")
}

func TestHtmlSimpleFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()

	SimpleHtmlFile(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}

// template directory untuk menangkap semua file yang ada dii dallaam directory
func SimpleHtmlFileDirectory(write http.ResponseWriter, request *http.Request) {
	// di sini menangkap semua file yang ada di dalam ./templates/*.gohtml menggunakan ParseGlob
	// jika ada file di dalam templates dengan akhiran .gohtml maka dia akan di ambil oleh ParseGlob
	htmlFile := template.Must(template.ParseGlob("./templates/*.gohtml"))
	// untuk mengakses file yang sudah di ambil kita bissa menggunakan dengan memanggil nama file yang sama dengan yang ingin kita gunakan
	htmlFile.ExecuteTemplate(write, "simple.gohtml", "ini html file directory")
}

func TestHtmlSimpleFileDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()

	SimpleHtmlFileDirectory(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}

// golang embed ini lebih di rekomendasikan
//
//go:embed templates/*.gohtml
var resorces embed.FS

func HtmlSimpleFileEmbed(write http.ResponseWriter, request *http.Request) {
	// untuk menggunakan embed di file ini kita kita gunakan ParseFS(var_embednya, "nama_filenya")
	// ParseFS(var_embednya, "templates/*.gohtml") ini akan memparse semua embed yang ada di dalam folder templates dan semua file dngan format .gohtml
	templateTextEmbed := template.Must(template.ParseFS(resorces, "templates/*.gohtml"))
	templateTextEmbed.ExecuteTemplate(write, "simple.gohtml", "ini menggunakan file embed")
}

func TestHtmlSimpleFileEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()
	HtmlSimpleFileEmbed(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}
