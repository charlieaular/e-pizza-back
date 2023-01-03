package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Image       string
	Description string
	Weight      string
	Price       int
	CategoryID  int
	Category    Category
}
