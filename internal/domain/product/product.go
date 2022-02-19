package product

import "fmt"

type Product struct {
	Name string
}

func New(name string) Product {
	return Product{Name: name}
}

func (p Product) String() string {
	return fmt.Sprintf("Product{Name: %s}", p.Name)
}
