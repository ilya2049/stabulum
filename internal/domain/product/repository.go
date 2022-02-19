package product

type Repository interface {
	Add(Product) error
}
