package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// jika kita menggunakan CrossSiteCripting auto compile maka kode kode html/css/js akan di compile ke kode yang ada di golang secara otomatis
func TemplateCrossSiteScripting(w http.ResponseWriter, r *http.Request) {
	templateXss := template.Must(template.ParseFiles("./templates/postXss.gohtml"))
	templateXss.ExecuteTemplate(w, "postXss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>INI adalah halaman body</p>",
	})
}

func TestTemplateCrossSiteScripting(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	reco := httptest.NewRecorder()

	TemplateCrossSiteScripting(reco, req)

	body, _ := io.ReadAll(reco.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateCrossSiteScriptingWeb(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:3030",
		Handler: http.HandlerFunc(TemplateCrossSiteScripting),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// XSS Di matikan
func TemplateCrossSiteScriptingDisable(w http.ResponseWriter, r *http.Request) {
	templateXss := template.Must(template.ParseFiles("./templates/postXss.gohtml"))
	templateXss.ExecuteTemplate(w, "postXss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTMLEscaper("<p>INI adalah halaman body</p><script>alert(`hello`)</script>"),
	})
}

func TestTemplateCrossSiteScriptingDisable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	reco := httptest.NewRecorder()

	TemplateCrossSiteScriptingDisable(reco, req)

	body, _ := io.ReadAll(reco.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateCrossSiteScriptingWebDisable(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateCrossSiteScriptingDisable),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// XSS Di matikan
func TemplateCrossSiteScriptingUser(w http.ResponseWriter, r *http.Request) {
	templateXss := template.Must(template.ParseFiles("./templates/postXss.gohtml"))
	templateXss.ExecuteTemplate(w, "postXss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTMLEscaper(r.URL.Query().Get("body")),
	})
}

func TestTemplateCrossSiteScriptingUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	reco := httptest.NewRecorder()

	TemplateCrossSiteScriptingUser(reco, req)

	body, _ := io.ReadAll(reco.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateCrossSiteScriptingWebUser(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateCrossSiteScriptingUser),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
