package queries

import (
	"context"
	"fmt"
	"stabulum/internal/app/queries"
	"stabulum/internal/common/logger"
	"stabulum/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type Querier struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewQuerier(db *gorm.DB, logger logger.Logger) *Querier {
	return &Querier{
		db:     db,
		logger: logger,
	}
}

func (*Querier) buildFindProductsQuery(query queries.ProductListQuery) string {
	filter := ""
	if query.Name != "" {
		filter = fmt.Sprintf(" WHERE name = '%s'", query.Name)
	}

	return fmt.Sprintf("SELECT id, name FROM products%s", filter)
}

func (q *Querier) FindProducts(ctx context.Context, query queries.ProductListQuery,
) (queries.ProductList, error) {
	productList := queries.ProductList{}

	err := q.db.Raw(q.buildFindProductsQuery(query)).
		Scan(&productList).
		Error

	if err != nil {
		return queries.ProductList{},
			fmt.Errorf("%s: failed to retrieve a product list: %w", postgres.Component, err)
	}

	return productList, nil
}
