package app

import (
	"database/sql"
	"restfullapi_news/helper"
	"time"
)

func ConnectToDB() *sql.DB {
	DB, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_project?parseTime=true")
	helper.IfLogingErr(err, "error terjadi di dalam app db_connection sql Open")

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(60 * time.Minute)
	DB.SetConnMaxLifetime(5 * time.Minute)

	return DB
}
