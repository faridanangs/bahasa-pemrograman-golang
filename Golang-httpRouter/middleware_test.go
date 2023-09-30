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

type LogMiddleware struct {
	http.Handler
}

func (middle *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Middleware execute")
	middle.Handler.ServeHTTP(w, r)
}

func TestHttpRouterMiddleware(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	logMiddleware := &LogMiddleware{Handler: router}

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recod := httptest.NewRecorder()

	logMiddleware.ServeHTTP(recod, req)

	body, _ := io.ReadAll(recod.Result().Body)

	assert.Equal(t, "Middleware", string(body))

}
