package product

import (
	"stabulum/internal/domain/product"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ProductView struct {
	Name string `json:"name"`
}

func (p ProductView) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(5, 50)),
	)
}

func (p ProductView) AsProduct() product.Product {
	return product.New(p.Name)
}
