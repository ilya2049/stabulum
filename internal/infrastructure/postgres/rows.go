package postgres

import "database/sql"

func CloseRows(rows *sql.Rows, handleError func(error)) {
	if err := rows.Close(); err != nil {
		handleError(err)
	}
}

func HandleRowsError(rows *sql.Rows, handleError func(error)) {
	if err := rows.Err(); err != nil {
		handleError(err)
	}
}
