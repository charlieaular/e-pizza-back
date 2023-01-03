package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"e-pizza-backend/src/services"
	"e-pizza-backend/src/shared"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (h *productHandler) GetProductsByCategory(c *gin.Context) {

	categoryId, err := strconv.Atoi(c.Param("category"))

	shared.HandleError(c, err)

	products, err := h.productService.GetProductsByCategory(categoryId)

	shared.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})
}

func (h *productHandler) GetProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	shared.HandleError(c, err)

	product, err := h.productService.GetById(id)

	shared.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})
}
