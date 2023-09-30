package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, r.Body)
	// 	fmt.Fprintln(w, r.Method)
	// 	fmt.Fprintln(w, r.Close)
	// 	fmt.Fprintln(w, r.ContentLength)
	// 	fmt.Fprintln(w, r.Form)
	// 	fmt.Fprintln(w, r.Header)
	// 	fmt.Fprintln(w, r.MultipartForm)
	// 	fmt.Fprintln(w, r.PostForm)
	// 	fmt.Fprintln(w, r.Proto)
	// 	fmt.Fprintln(w, r.RemoteAddr)
	// 	fmt.Fprintln(w, r.RequestURI)
	// 	fmt.Fprintln(w, r.Response)
	// 	fmt.Fprintln(w, r.URL)
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprintln(w, "<h1>Hello dunia</h1>")
		fmt.Fprintln(w, "<h4>Hello Body</h4>", r.Body)
		fmt.Fprintln(w, "<h4>Hello method</h4>", r.Method)
		fmt.Fprintln(w, "<h4>Hello close</h4>", r.Close)
		fmt.Fprintln(w, "<h4>Hello ContentLength</h4>", r.ContentLength)
		fmt.Fprintln(w, "<h4>Hello Form</h4>", r.Form)
		fmt.Fprintln(w, "<h4>Hello Header</h4>", r.Header)
		fmt.Fprintln(w, "<h4>Hello MultipartForm</h4>", r.MultipartForm)
		fmt.Fprintln(w, "<h4>Hello PostForm</h4>", r.PostForm)
		fmt.Fprintln(w, "<h4>Hello Proto</h4>", r.Proto)
		fmt.Fprintln(w, "<h4>Hello RemoteAddr</h4>", r.RemoteAddr)
		fmt.Fprintln(w, "<h4>Hello RequestURI</h4>", r.RequestURI)
		fmt.Fprintln(w, "<h4>Hello Response</h4>", r.Response)
		fmt.Fprintln(w, "<h4>Hello URL</h4>", r.URL)
		fmt.Fprintln(w, "<h4>Hello TLS</h4>", r.TLS)
		fmt.Fprintln(w, "<h4>Hello Trailer</h4>", r.Trailer)
		fmt.Fprintln(w, "<h4>Hello TransferEncoding</h4>", r.TransferEncoding)

		fmt.Fprintln(w, "<h4>Hello URL ForceQuery</h4>", r.URL.ForceQuery)
		fmt.Fprintln(w, "<h4>Hello URL Fragment</h4>", r.URL.Fragment)
		fmt.Fprintln(w, "<h4>Hello URL Host</h4>", r.URL.Host)
		fmt.Fprintln(w, "<h4>Hello URL Hostname</h4>", r.URL.Hostname())
		fmt.Fprintln(w, "<h4>Hello URL OmitHost</h4>", r.URL.OmitHost)
		fmt.Fprintln(w, "<h4>Hello URL Opaque</h4>", r.URL.Opaque)
		fmt.Fprintln(w, "<h4>Hello URL Path</h4>", r.URL.Path)
		fmt.Fprintln(w, "<h4>Hello URL RawPath</h4>", r.URL.RawPath)
		fmt.Fprintln(w, "<h4>Hello URL RawFragment</h4>", r.URL.RawFragment)
		fmt.Fprintln(w, "<h4>Hello URL RawQuery</h4>", r.URL.RawQuery)
		fmt.Fprintln(w, "<h4>Hello URL Scheme</h4>", r.URL.Scheme)
		fmt.Fprintln(w, "<h4>Hello URL User</h4>", r.URL.User)
		fmt.Fprintln(w, "<h4>Hello URL Query().Get())</h4>", r.URL.Query().Get("name"))
		fmt.Fprintln(w, "<h4>Hello URL User.Username()</h4>", r.URL.User.Username())
		fmt.Fprintln(w, "<h4>Hello URL Query().Encode()</h4>", r.URL.Query().Encode())
		fmt.Fprintln(w, "<h4>Hello URL Query().Encode()</h4>", r.URL.Query().Has("name"))
		fmt.Fprintln(w, "<h4>Hello URL Port()</h4>", r.URL.Port())
		fmt.Fprintln(w, "<h4>Hello URL Redacted</h4>", r.URL.Redacted())
		fmt.Fprintln(w, "<h4>Hello URL EscapedFragment</h4>", r.URL.EscapedFragment())
		fmt.Fprintln(w, "<h4>Hello URL RequestURI</h4>", r.URL.RequestURI())
		fmt.Fprintln(w, "<h4>Hello URL String</h4>", r.URL.String())

	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
