package product

import "stabulum/internal/domain/product"

type ProductView struct {
	Name string
}

func (p ProductView) AsProduct() product.Product {
	return product.New(p.Name)
}
