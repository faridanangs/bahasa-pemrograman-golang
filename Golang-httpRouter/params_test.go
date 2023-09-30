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

func TestHttpRouterParams(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// kita gunakan p.ByName("id") untuk emnangkap id dari product yang di kirim oleh user di url
		text := "Product " + params.ByName("id")
		fmt.Fprint(w, text)
	})

	// nama dari NewRequest harus sesuai dengan nama method yang kita berikan seperti GET dll
	req := httptest.NewRequest("GET", "http://localhost:3000/product/2", nil)
	recod := httptest.NewRecorder()

	// kemudian kita gunakan serverhttp untuk melakukan unit test
	router.ServeHTTP(recod, req)

	response := recod.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(body))

}
