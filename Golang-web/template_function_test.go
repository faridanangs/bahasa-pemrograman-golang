package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type Sapa struct {
	Name string
}

func (sapa Sapa) SapaDia(name string) string {
	return "Hello " + name + " My name is " + sapa.Name
}

func TemplateFuntion(write http.ResponseWriter, request *http.Request) {
	// func SapaDia ini dia mengambil di dalam strct saya di dalam ExecuteTemplate
	// jika kita tidak mengirim struct nya maka function ini tidak di ketahui
	templateFunc := template.Must(template.New("Function").Parse(`{{.SapaDia "Anangs"}}`))
	// templateFunc.ExecuteTemplate(write, "Function", `Sapa{Name: "Diana"}`) jika seperti ini dia akan kosong karna SapaDia ini tidak di ketahui
	templateFunc.ExecuteTemplate(write, "Function", Sapa{Name: "Diana"})
}

func TestTemplateFuntion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateFuntion(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}

// template global function yang sudah di sediakan oleh golang
func TemplateFuntionGlobal(write http.ResponseWriter, request *http.Request) {
	// func SapaDia ini dia mengambil di dalam strct saya di dalam ExecuteTemplate
	// jika kita tidak mengirim struct nya maka function ini tidak di ketahui
	templateFunc := template.Must(template.New("Function").Parse(`{{len "Anangs"}}`))
	// templateFunc.ExecuteTemplate(write, "Function", `Sapa{Name: "Diana"}`) jika seperti ini dia akan kosong karna SapaDia ini tidak di ketahui
	templateFunc.ExecuteTemplate(write, "Function", Sapa{Name: "Diana"})
}

func TestTemplateFuntionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateFuntionGlobal(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}

// kita juga bisa membuat global function di golang
func TemplateCreateFuntionGlobal(write http.ResponseWriter, request *http.Request) {

	templateFunc := template.New("Function")
	// untuk membuat global function di golang kita hanya menggunakan templateFunc.Funcs()
	templateFunc.Funcs(map[string]interface{}{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
	})
	// setelah jita membuat global funct kita baru bisa emparsenya dan memasukan nama global funcnya
	template.Must(templateFunc.Parse(`{{upper "farid anang s"}}`))
	templateFunc.ExecuteTemplate(write, "Function", Sapa{Name: "Jodu"})
}

func TestTemplateCreateFuntionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateCreateFuntionGlobal(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}

// kita juga bisa membuat global function di golang
func TemplateCreateFuntionGlobalFipeLine(write http.ResponseWriter, request *http.Request) {

	templateFunc := template.New("Function")
	// untuk membuat global function di golang kita hanya menggunakan templateFunc.Funcs()
	templateFunc.Funcs(map[string]interface{}{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
		"sayHello": func(name string) string {
			return "hello " + name
		},
	})

	// fipeline adalah kondii di mana kita membagi funtion kita contohnya seperti ini {{sayHello "farid anang s" | upper}} pertama" funtion sayHello menerima
	// parameter "farid anang s" setelah di eksekusi maka nilai yang di return func sayHello akan di teruskan ke function upper dengan menggunakan |
	template.Must(templateFunc.Parse(`{{sayHello "farid anang s" | upper}}`))
	templateFunc.ExecuteTemplate(write, "Function", Sapa{Name: "Jodu"})
}

func TestTemplateCreateFuntionGlobalFipeLine(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recod := httptest.NewRecorder()

	TemplateCreateFuntionGlobalFipeLine(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)
	fmt.Println(string(body))
}
