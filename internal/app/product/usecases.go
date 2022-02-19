package product

import "stabulum/internal/domain/product"

type Usecases struct {
	productRepostiory product.Repository
}

func NewUsecases(productRepostiory product.Repository) *Usecases {
	return &Usecases{
		productRepostiory: productRepostiory,
	}
}

func (u *Usecases) Create(p product.Product) error {
	return u.productRepostiory.Add(p)
}
