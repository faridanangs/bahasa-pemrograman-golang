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

//go:embed templates/*.gohtml
var resorcesDinamis embed.FS

// menggunakan map[string]string
func TemplateDataDinamisMap(write http.ResponseWriter, request *http.Request) {
	templHtmlDinamis := template.Must(template.ParseFS(resorcesDinamis, "templates/*.gohtml"))
	// untuk membuat template data dinamis kita bisa menggunakan struct atau map tempat memasukan nilai atau datanya kemudian cara memanggilnya
	// kita cukup panggil keynya di file html contoh: {{.Name}} jika bersarang ckup use {{.Addres.Stret}}
	templHtmlDinamis.ExecuteTemplate(write, "name.gohtml", map[string]interface{}{
		"Title": "Tempalte dinamis",
		"Name":  "farid anang s",
	})
}

func TestTemplateDataDinamis(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()
	TemplateDataDinamisMap(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}

type Alamat struct {
	Stret string
}
type Data struct {
	Title, Name string
	Address     Alamat
}

// menggunakan struct
func TemplateDataDinamisStruct(write http.ResponseWriter, request *http.Request) {
	templHtmlDinamis := template.Must(template.ParseFS(resorcesDinamis, "templates/*.gohtml"))
	// untuk membuat template data dinamis kita bisa menggunakan struct atau map tempat memasukan nilai atau datanya kemudian cara memanggilnya
	// kita cukup panggil keynya di file html contoh: {{.Name}} jika bersarang ckup use {{.Addres.Stret}}
	templHtmlDinamis.ExecuteTemplate(write, "name.gohtml", Data{
		Name:  "Farid Anang S",
		Title: "Ini menggunakan Struct",
		// begini cara membuat sruct bersarang
		Address: Alamat{
			Stret: "sesela kebon lauk",
		},
	})
}
func TestTemplateDataDinamisStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()
	TemplateDataDinamisStruct(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}
