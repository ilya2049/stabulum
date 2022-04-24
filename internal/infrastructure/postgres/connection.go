package postgres

import (
	"database/sql"
	"fmt"
	"stabulum/internal/common/logger"
	"stabulum/internal/pkg/connection"

	// Postgres driver.
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const Component = "postgres"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Connection struct {
	SQLDB  *sql.DB
	GormDB *gorm.DB
}

func NewConnection(cfg Config, logger logger.Logger) (Connection, connection.Close, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	gormDB, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return Connection{}, nil, fmt.Errorf("%s: failed to create a postgres connection: %w", Component, err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return Connection{}, nil, fmt.Errorf("%s: failed to separate sql db from gorm db: %w", Component, err)
	}

	logger.Println(Component + ": connection established")

	return Connection{SQLDB: sqlDB, GormDB: gormDB}, func() {
		if err := sqlDB.Close(); err != nil {
			logger.Println(Component + ": failed to close postgres connection pull: " + err.Error())

			return
		}

		logger.Println(Component + ": connection closed")
	}, nil
}
