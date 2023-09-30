package golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	// untuk memudahka kita supaya bisa terhubung ke database kita bisa menggunakan
	// sql.Opne("driver/nama-db", "username:password(localhost:port)/nama-db")
	db, err := sql.Open("mysql", "root:root(localhost:3306)/db_golang")
	if err != nil {
		panic(err)
	}
	// kemudian kita harus menclose dbnya jika sudah tidak di guunakan supaya tidak terjadi crash
	defer db.Close()
}
