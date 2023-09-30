package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type CommentRepositoryIplm struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryIplm{DB: db}
}

// Memasukan data ke dalam database
func (repository *CommentRepositoryIplm) Insert(cntx context.Context, commnet entity.Comment) (entity.Comment, error) {
	script := "insert into comments(email, comment) values(?, ?)"
	// jika kita gunakan ExecContext maka kita akan mengirim perintah ke database
	result, err := repository.DB.ExecContext(cntx, script, commnet.Email, commnet.Comment)

	if err != nil {
		panic(err)
	}

	IncId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	commnet.Id = int32(IncId)
	return commnet, nil
}

func (repository *CommentRepositoryIplm) FindById(cntx context.Context, id int32) (entity.Comment, error) {
	script := "select id, email, comment from comments where id = ? limit 1"

	// jika kita gunakan QueryContext maka kita akan menerima data dari database
	rows, err := repository.DB.QueryContext(cntx, script, id)
	defer rows.Close()

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *CommentRepositoryIplm) FindByAll(cntx context.Context) ([]entity.Comment, error) {
	script := "select * from comments"

	rows, err := repository.DB.QueryContext(cntx, script)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}

// userrr
type UserRepositoryIplm struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryIplm{DB: db}
}

func (repo *UserRepositoryIplm) InsertUser(cntx context.Context, user entity.User) (entity.User, error) {
	script := "insert into customer(id, name, email, balance, rating, birth_date) value(?,?,?,?,?,?)"

	_, err := repo.DB.ExecContext(cntx, script, user.Id, user.Name, user.Email, user.Balance, user.Rating, user.BirthDay)
	if err != nil {
		panic(err)
	}

	return user, nil
}

func (repo *UserRepositoryIplm) FindUserAll(cntx context.Context) ([]entity.User, error) {
	script := "SELECT id, name, email, balance, rating, birth_date FROM customer"
	rows, err := repo.DB.QueryContext(cntx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Balance, &user.Rating, &user.BirthDay)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
