package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TempalateLayout(write http.ResponseWriter, request *http.Request) {
	htmlLayout := template.Must(template.ParseFiles("./templates/footer.gohtml", "./templates/header.gohtml", "./templates/layout.gohtml"))
	htmlLayout.ExecuteTemplate(write, "layout", map[string]interface{}{
		"Title": "template action",
		"Name":  "FaridAnangS",
		"Address": map[string]string{
			"Jalan": "suiman 01 323",
			"Dusun": "sesela kebon lauk",
		}})
}

func TestTempalateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TempalateLayout(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}
