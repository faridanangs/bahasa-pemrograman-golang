package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// if di dalam template
func TemplateActionIf(write http.ResponseWriter, request *http.Request) {
	htmlFile := template.Must(template.ParseFiles("./templates/if.gohtml"))
	htmlFile.ExecuteTemplate(write, "if.gohtml", map[string]interface{}{
		"Title": "template action",
		"Name":  "Farid anangs ",
	})
}

func TestTemplateActionIf(t *testing.T) {
	// req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	// recod := httptest.NewRecorder()

	// TemplateActionIf(recod, req)
	// body, _ := io.ReadAll(recod.Result().Body)
	// fmt.Println(string(body))
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateActionIf)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// oprator logika di dalam template
func TemplateActionOprator(write http.ResponseWriter, request *http.Request) {
	htmlFile := template.Must(template.ParseFiles("./templates/oprator.gohtml"))
	htmlFile.ExecuteTemplate(write, "oprator.gohtml", map[string]interface{}{
		"Title": "template action",
		"Value": 40,
	})
}

func TestTemplateActionOprator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateActionOprator(recod, req)
	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}

// range/iterasi di dalam template
func TemplateActionRange(write http.ResponseWriter, request *http.Request) {
	htmlFile := template.Must(template.ParseFiles("./templates/range.gohtml"))
	htmlFile.ExecuteTemplate(write, "range.gohtml", map[string]interface{}{
		"Title": "template action",
		"Hobbies": []string{
			"cooding", "makan", "running", "workout",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateActionRange(recod, req)
	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}

// with di gunakan untum me nested struct,map,slice,array yang bersarang di dalam template
func TemplateActionWith(write http.ResponseWriter, request *http.Request) {
	htmlFile := template.Must(template.ParseFiles("./templates/address.gohtml"))
	htmlFile.ExecuteTemplate(write, "address.gohtml", map[string]interface{}{
		"Title": "template action",
		"Address": map[string]string{
			"Jalan": "suiman 01 323",
			"Dusun": "sesela kebon lauk",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateActionWith(recod, req)
	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}
