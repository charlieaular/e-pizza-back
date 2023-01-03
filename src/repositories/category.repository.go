package repositories

import (
	"gorm.io/gorm"

	"e-pizza-backend/src/models"
)

type CategoryRepo interface {
	GetAll() ([]models.Category, error)
}

type categoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{DB: db}
}

func (categoryRepo *categoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category

	result := categoryRepo.DB.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}
