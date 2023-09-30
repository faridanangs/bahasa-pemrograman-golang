package golanghttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHttpRouterNotFound(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	// jika url tidak ada maka akan terjadi notfound
	// untuk menangkap notfound kita bisa menggunakan router.NotFound = http.Handler dan ini yang akan di jalankan
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "halaman gak ada")
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recod := httptest.NewRecorder()

	router.ServeHTTP(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)

	assert.Equal(t, "halaman gak ada", string(body))

}
