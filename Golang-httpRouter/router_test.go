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

func TestHttpRouterLib(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	// httprouter ini memiliki banyak sekali method yang tidak di mliki oleh router bawaan golang
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	// nama dari NewRequest harus sesuai dengan nama method yang kita berikan seperti GET dll
	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recod := httptest.NewRecorder()

	// kemudian kita gunakan serverhttp untuk melakukan unit test
	router.ServeHTTP(recod, req)

	response := recod.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(body))

}
