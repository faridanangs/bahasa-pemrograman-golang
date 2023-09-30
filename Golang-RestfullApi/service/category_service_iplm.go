package service

import (
	"context"
	"database/sql"
	"golangapi/exception"
	"golangapi/helper"
	"golangapi/model/domain"
	"golangapi/model/web"
	"golangapi/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceIplm struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceIplm{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceIplm) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// melakukan validate kepada request yang di kirim
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// kita mendapatkan nilai tx di dalam db.begin()
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Name: request.Name}
	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceIplm) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// ambil data yang mau di update
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NotFound(err.Error()))
	}
	category = domain.Category{Name: request.Name, Id: request.Id}
	service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceIplm) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NotFound(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category)
}
func (service *CategoryServiceIplm) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NotFound(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceIplm) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
