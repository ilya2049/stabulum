package queries

import (
	"context"
	"database/sql"
	"fmt"
	"stabulum/internal/app/queries"
	"stabulum/internal/common/logger"
	"stabulum/internal/infrastructure/postgres"

	"github.com/huandu/go-sqlbuilder"
)

type Querier struct {
	db     *sql.DB
	logger logger.Logger
}

func NewQuerier(db *sql.DB, logger logger.Logger) *Querier {
	return &Querier{
		db:     db,
		logger: logger,
	}
}

func (q *Querier) FindProducts(ctx context.Context, query queries.ProductListQuery,
) (queries.ProductList, error) {
	sqlQuery := sqlbuilder.Select("id", "name").From("products")
	if query.Name != "" {
		sqlQuery.Where("name = '" + query.Name + "'")
	}

	rows, err := q.db.QueryContext(ctx, sqlQuery.String())
	if err != nil {
		return queries.ProductList{}, fmt.Errorf(
			"%s: failed to retrieve products: %w", postgres.Component, err,
		)
	}

	defer postgres.HandleRowsError(rows, func(err error) {
		q.logger.Println(
			fmt.Sprintf("%s: rows error after retrieving products: %s", postgres.Component, err.Error()),
		)
	})

	defer postgres.CloseRows(rows, func(err error) {
		q.logger.Println(
			fmt.Sprintf("%s: failed to close rows after retrieving products: %s", postgres.Component, err.Error()),
		)
	})

	productList := queries.ProductList{}
	var productListItem queries.ProductListItem

	for rows.Next() {
		err = rows.Scan(
			&productListItem.ID,
			&productListItem.Name,
		)

		if err != nil {
			return queries.ProductList{}, fmt.Errorf("%s: failed to scan a product: %w",
				postgres.Component, err,
			)
		}

		productList = append(productList, productListItem)
	}

	return productList, nil
}
