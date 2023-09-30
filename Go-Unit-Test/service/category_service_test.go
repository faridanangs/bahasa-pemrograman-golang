package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	require.Nil(t, category)
	require.NotNil(t, err)

}

func TestCategoryServiceFound(t *testing.T) {
	category := entity.Category{
		Name: "farid",
		Id:   "1",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	require.Nil(t, err)
	require.NotNil(t, result)

	require.Equal(t, category.Id, result.Id)
	require.Equal(t, category.Name, result.Name)

}
