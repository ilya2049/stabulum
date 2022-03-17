package product

import (
	"net/http"

	"stabulum/internal/app/product"
	"stabulum/internal/app/queries"
	"stabulum/internal/infrastructure/api"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	productUsecases product.Usecases
	produtcQuerier  queries.ProductQuerier
}

func NewHandler(
	productUsecases product.Usecases,
	produtcQuerier queries.ProductQuerier,
) *Handler {
	return &Handler{
		productUsecases: productUsecases,
		produtcQuerier:  produtcQuerier,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var productView ProductView
	if err := api.ShouldBindJSON(c, &productView); err != nil {
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

func (h *Handler) List(c *gin.Context) {
	productQuery := queries.ProductListQuery{
		Name: c.Query("name"),
	}

	productList, err := h.produtcQuerier.FindProducts(c.Request.Context(), productQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: err.Error()})

		return
	}

	c.JSON(http.StatusOK, productList)
}
