package database

import (
	"database/sql"
	"fmt"
	"linktree/config"

	_ "github.com/lib/pq"
)

func ConnectToPostgres(cfg *config.AppConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		cfg.PostgresDatabase.Host,
		cfg.PostgresDatabase.Port,
		cfg.PostgresDatabase.Username,
		cfg.PostgresDatabase.Password,
		cfg.PostgresDatabase.Name,
		cfg.PostgresDatabase.SslMode,
		cfg.PostgresDatabase.Timezone,
	)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.PostgresDatabase.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.PostgresDatabase.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(cfg.PostgresDatabase.ConnectionMaxLifetime)

	return sqlDB, nil
}
