package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Status     string
	TotalPrice int
}
