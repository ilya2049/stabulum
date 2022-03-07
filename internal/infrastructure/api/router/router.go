package router

import (
	"stabulum/internal/infrastructure/api/router/product"

	"github.com/gin-gonic/gin"
)

func New(productHandler *product.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/product", productHandler.Create)

	return r
}
