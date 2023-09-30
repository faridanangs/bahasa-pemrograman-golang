package repository

import (
	"context"
	"golang-database/entity"
)

type CommentRepository interface {
	Insert(cntx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(cntx context.Context, id int32) (entity.Comment, error)
	FindByAll(cntx context.Context) ([]entity.Comment, error)
}

type UserRepository interface {
	InsertUser(cntx context.Context, user entity.User) (entity.User, error)
	FindUserAll(cntx context.Context) ([]entity.User, error)
}
