package queries

import "context"

type ProductListQuery struct {
	Name string
}

type ProductList []ProductListItem

type ProductListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductQuerier interface {
	FindProducts(context.Context, ProductListQuery) (ProductList, error)
}
