package postgres

import "stabulum/internal/pkg/connection"

type Connection struct {
}

func NewConnection() (*Connection, connection.Close, error) {
	return &Connection{}, func() {}, nil
}
