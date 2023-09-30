package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// supaya tidak terjadi error ketika kita mengirim time.Time ke dalam database kita tambahkan
	// ?parseTime=true di akhir nama databasenya
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_golang?parseTime=true")
	if err != nil {
		return nil
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
