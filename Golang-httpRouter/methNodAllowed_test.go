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

func TestHttpRouterMethNodAllowed(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	// untuk menangkap method yang tidak sesuai kita bisa menggunakan router.MethodNotAllowed = http.Handler
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Post")
	})

	req := httptest.NewRequest("PUT", "http://localhost:8080/", nil)
	recod := httptest.NewRecorder()

	router.ServeHTTP(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)

	assert.Equal(t, "Gak Boleh", string(body))

}
