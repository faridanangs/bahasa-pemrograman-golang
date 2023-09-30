package repository

import (
	"context"
	"fmt"
	go_db "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_db.GetConnection())

	cntx := context.Background()
	comment := entity.Comment{
		Email:   "Faridanangs@gmail.com",
		Comment: "Hello whatUp man",
	}

	result, err := commentRepository.Insert(cntx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_db.GetConnection())
	cntx := context.Background()

	for i := 1; i < 5; i++ {
		result, err := commentRepository.FindById(cntx, int32(i))
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_db.GetConnection())
	cntx := context.Background()

	result, err := commentRepository.FindByAll(cntx)
	if err != nil {
		panic(err)
	}

	for _, data := range result {
		fmt.Println(data)
	}
}

// User

func TestInsertUser(t *testing.T) {
	userRepository := NewUserRepository(go_db.GetConnection())

	cntx := context.Background()

	userData := entity.User{
		Id:      "puspu",
		Name:    "Farid usiidnubD",
		Balance: 10000000,
		Rating:  10.5,
	}
	result, err := userRepository.InsertUser(cntx, userData)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindUserAll(t *testing.T) {
	userRepository := NewUserRepository(go_db.GetConnection())
	cntx := context.Background()
	result, err := userRepository.FindUserAll(cntx)
	if err != nil {
		panic(err)
	}

	for id, data := range result {
		fmt.Println(id, data)
	}
}
