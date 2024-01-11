package repository

import "github.com/MuhammadIbraAlfathar/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
