package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e-pizza-backend/src/services"
	"e-pizza-backend/src/shared"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) categoryHandler {
	return categoryHandler{categoryService: categoryService}
}

func (h *categoryHandler) GetAllCategories(c *gin.Context) {

	categories, err := h.categoryService.GetAll()

	shared.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
	})
}
