package entity

import (
	"database/sql"
)

type Comment struct {
	Id             int32
	Email, Comment string
}

type User struct {
	Id, Name string
	Email    sql.NullString
	Balance  int64
	Rating   float64
	BirthDay sql.NullTime
}
