package repositories

import (
	"gorm.io/gorm"

	"e-pizza-backend/src/models"
)

type ProductRepo interface {
	GetProductsByCategory(int) ([]models.Product, error)
	GetById(int) (models.Product, error)
	GetByIds([]int) ([]models.Product, error)
}

type productRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{DB: db}
}

func (productRepo *productRepo) GetProductsByCategory(category int) ([]models.Product, error) {
	var products []models.Product

	result := productRepo.DB.Model(models.Product{}).Preload("Category").Where("category_id = ?", category).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

func (productRepo *productRepo) GetById(id int) (models.Product, error) {
	var product models.Product
	result := productRepo.DB.Model(models.Product{}).First(&product, id)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

func (productRepo *productRepo) GetByIds(ids []int) ([]models.Product, error) {
	var products []models.Product
	result := productRepo.DB.Model(models.Product{}).Find(&products, ids)

	if result.Error != nil {
		return []models.Product{}, result.Error
	}

	return products, nil
}
