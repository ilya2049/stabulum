package postgres

import (
	"database/sql"
	"fmt"
	"stabulum/internal/common/logger"
	"stabulum/internal/pkg/connection"

	// Postgres driver.
	_ "github.com/lib/pq"
)

const Component = "postgres"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConnection(cfg Config, logger logger.Logger) (*sql.DB, connection.Close, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	connection, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: failed to create a postgres connection: %w", Component, err)
	}

	if err := connection.Ping(); err != nil {
		return nil, nil, fmt.Errorf("%s: failed to connect to postgres: %w", Component, err)
	}

	logger.Println(Component + ": connection established")

	return connection, func() {
		if err := connection.Close(); err != nil {
			logger.Println(Component + ": failed to close postgres connection pull: " + err.Error())

			return
		}

		logger.Println(Component + ": connection closed")
	}, nil
}
