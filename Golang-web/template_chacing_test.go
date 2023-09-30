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

type AddressT struct {
	Stret string
}
type Datas struct {
	Name    string
	Address AddressT
	Title   string
}

// supaya template kita hanya di eksekusi sekali saja maka kita bisa menggunakan embed dan memanggilnya di global
// bukan memanggilnya di dalam funtion yang di tuju itu menyebabkan pepmengkakan karna setian func itu di panggil
// maka templatenya juga ikut terpanggil

//go:embed templates/*.gohtml
var templates embed.FS

// var templateChacing = template.Must(template.ParseFS(templates, "templates/*.gohtml"))
var templateChacing = template.Must(template.ParseGlob("./templates/*.gohtml"))

func TemplateChacing(w http.ResponseWriter, r *http.Request) {
	templateChacing.ExecuteTemplate(w, "name.gohtml", Datas{
		Title:   "Chacing template",
		Name:    "farid anang s",
		Address: AddressT{Stret: "sesela"},
	})
}

func TestTemplateChacing(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recode := httptest.NewRecorder()
	TemplateChacing(recode, request)

	body, _ := io.ReadAll(recode.Result().Body)
	fmt.Println(string(body))
}
