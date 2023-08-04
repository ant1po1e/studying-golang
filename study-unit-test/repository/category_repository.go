package repository

import "study-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}