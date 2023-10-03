package test

import (
	"database/sql"
	"restfullapi_news/helper"
	"time"
)

func setUpDB() *sql.DB {
	DB, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_project?parseTime=true")
	helper.IfLogingErr(err, "error terjadi di dalam test setupDB sql Open")

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
