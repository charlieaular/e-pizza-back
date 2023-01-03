package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-pizza-backend/src/handlers"
	"e-pizza-backend/src/repositories"
	"e-pizza-backend/src/services"
)

func RegisteProductRoutes(router *gin.Engine, db *gorm.DB) {

	productRepo := repositories.NewProductRepo(db)
	productService := services.ProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRoutes := router.Group("products")
	{
		categoryRoutes.GET("/:id", productHandler.GetProductById)
		categoryRoutes.GET("/category/:category", productHandler.GetProductsByCategory)

	}

}
