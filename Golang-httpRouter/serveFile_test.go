package golanghttprouter

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestHttpRouterServeFile(t *testing.T) {
	// kita buat dulu httprouternya
	router := httprouter.New()

	dir, _ := fs.Sub(resources, "resources")

	// kita menggunakan ServeFile untuk membaca text yang ada di dalam folder dan juga
	// kita harus menggunakan *filepath di akhir url servefilenya supaya tidak terjadi panis
	router.ServeFiles("/files/*filepath", http.FS(dir))

	// nama dari NewRequest harus sesuai dengan nama method yang kita berikan seperti GET dll
	req := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recod := httptest.NewRecorder()

	// kemudian kita gunakan serverhttp untuk melakukan unit test
	router.ServeHTTP(recod, req)

	response := recod.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(body))

}
