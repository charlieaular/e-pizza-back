package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-pizza-backend/src/handlers"
	"e-pizza-backend/src/repositories"
	"e-pizza-backend/src/services"
)

func RegisteCategoryRoutes(router *gin.Engine, db *gorm.DB) {

	categoryRepo := repositories.NewCategoryRepo(db)
	categoryService := services.CategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	categoryRoutes := router.Group("category")
	{
		categoryRoutes.GET("", categoryHandler.GetAllCategories)

	}

}
