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

func TestHttpRouterPanicHandler(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	// untuk menangkap panic kita bisa menggunakan router.PanicHandler dan di sini panicnya akan di simpan di dalam interpace kosong
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Error :", i)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recod := httptest.NewRecorder()

	router.ServeHTTP(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)

	assert.Equal(t, "Error :Ups", string(body))

}
