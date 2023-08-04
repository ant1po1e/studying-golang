package service

import (
	"study-unit-test/entity"
	"study-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(test *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(test, category)
	assert.NotNil(test, err)
}

func TestCategoryServie_GetSuccess(test *testing.T) {
	category := entity.Category {
		ID: "1",
		Name: "Something",	
	}
	
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")
	assert.Nil(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, category.ID, result.ID)
	assert.Equal(test, category.Name, result.Name)
}