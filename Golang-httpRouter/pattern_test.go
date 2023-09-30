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

// pattern name
func PatternNamed(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	text := "Product " + params.ByName("id") + " Item " + params.ByName("item")
	fmt.Fprint(w, text)
}

// pattern all
func PatternAllNamed(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	text := "Image: " + params.ByName("image")
	fmt.Fprint(w, text)
}

func TestHttpRouterPatternNamedParams(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	// router.GET("/product/:id/item/:item", PatternNamed)

	// untuk menangkap semua url kita bisa gunakan tanda * seperti ini /images/*image
	router.GET("/images/*image", PatternAllNamed)

	// nama dari NewRequest harus sesuai dengan nama method yang kita berikan seperti GET dll
	req := httptest.NewRequest("GET", "http://localhost:3000/images/resorces/anangs.png", nil)
	recod := httptest.NewRecorder()

	// kemudian kita gunakan serverhttp untuk melakukan unit test
	router.ServeHTTP(recod, req)

	response := recod.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image: /resorces/anangs.png", string(body))

}
