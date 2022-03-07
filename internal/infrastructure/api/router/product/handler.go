package product

import (
	"net/http"

	"stabulum/internal/app/product"
	"stabulum/internal/infrastructure/api"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	productUsecases product.Usecases
}

func NewHandler(productUsecases product.Usecases) *Handler {
	return &Handler{
		productUsecases: productUsecases,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var productView ProductView
	if err := c.ShouldBindJSON(&productView); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{Error: err.Error()})

		return
	}

	err := h.productUsecases.Create(c.Request.Context(), productView.AsProduct())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: err.Error()})

		return
	}

	c.Status(http.StatusCreated)
}
