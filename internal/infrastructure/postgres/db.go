package postgres

import (
	"context"
	"database/sql"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Tx interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}
