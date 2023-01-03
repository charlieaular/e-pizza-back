package services

import (
	"e-pizza-backend/src/models"
	"e-pizza-backend/src/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
}

type categoryService struct {
	categoryRepo repositories.CategoryRepo
}

func NewCategoryService(categoryRepo repositories.CategoryRepo) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (categoryService *categoryService) GetAll() ([]models.Category, error) {

	return categoryService.categoryRepo.GetAll()
}
