package main

import (
	"restfullapi_news/helper"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(t *testing.T) {
	server := WireInjectionBerita()
	if server != nil {
		err := server.ListenAndServe()
		helper.IfLogingErr(err, "error pada badian server ListenAndServe di test main")

	}
}
