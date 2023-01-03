package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-pizza-backend/src/models"
	requests "e-pizza-backend/src/requests/cart"
	"e-pizza-backend/src/services"
	"e-pizza-backend/src/shared"
)

type cartHandler struct {
	cartService        services.CartService
	productService     services.ProductService
	cartProductService services.CartProductService
}

func NewCartHandler(cartService services.CartService, productService services.ProductService, cartProductService services.CartProductService) cartHandler {
	return cartHandler{cartService: cartService, productService: productService, cartProductService: cartProductService}
}

func (h *cartHandler) GetCartProducts(c *gin.Context) {
	cartProducts, err := h.cartProductService.GetCartProducts()

	shared.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"CartProducts": cartProducts,
	})
}

func (h *cartHandler) AddProductToCart(c *gin.Context) {
	var addProductToCartRequest requests.AddProductToCartRequest

	err := c.ShouldBind(&addProductToCartRequest)

	shared.HandleError(c, err)

	product, err := h.productService.GetById(addProductToCartRequest.ProductID)

	shared.HandleError(c, err)

	currentCart, err := h.cartService.CurrentActiveCart()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		currentCart, err = h.cartService.CreateCart()
	}

	shared.HandleError(c, err)

	cartProduct := models.CartProduct{
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Weight:      product.Weight,
		Price:       product.Price,
		Quantity:    1,
		CategoryID:  product.CategoryID,
		CartID:      int(currentCart.ID),
	}

	newCartProduct, err := h.cartProductService.CreateCartProduct([]models.CartProduct{cartProduct})

	shared.HandleError(c, err)

	err1 := h.cartService.UpdateCartPrice()

	shared.HandleError(c, err1)

	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"newCartProduct": newCartProduct,
	})
}

func (h *cartHandler) DeleteCartProduct(c *gin.Context) {
	cartProductId, err := strconv.Atoi(c.Param("id"))

	shared.HandleError(c, err)

	status, err := h.cartProductService.DeleteCartProduct(cartProductId)

	shared.HandleError(c, err)

	cartProducts, err := h.cartProductService.GetCartProducts()

	shared.HandleError(c, err)

	if len(cartProducts) == 0 {
		err := h.cartService.DesactivateCart()

		shared.HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})

}

func (h *cartHandler) PayCart(c *gin.Context) {

	err := h.cartService.PayCart()

	shared.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
