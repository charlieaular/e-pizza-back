package models

import (
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	Name        string
	Image       string
	Description string
	Weight      string
	Price       int
	Quantity    int
	CategoryID  int
	CartID      int
	Category    Category
	Cart        Cart
}
