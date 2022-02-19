package product

import "context"

type Repository interface {
	Add(context.Context, Product) error
}
