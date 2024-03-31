package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	App struct {
		Name string
		Port string
		Env  string
	}

	Gin struct {
		Mode string
	}

	PostgresDatabase struct {
		Database
	}
}

type Database struct {
	Username              string
	Password              string
	Host                  string
	Name                  string
	Port                  uint64
	SslMode               string
	Timezone              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnectionMaxLifetime time.Duration
	MigrationPath         string
}

func NewConfig() (*AppConfig, error) {
	cfg := &AppConfig{}

	setAppConfig(cfg)
	setGinConfig(cfg)

	if err := setPostgresDbConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func setGinConfig(cfg *AppConfig) {
	cfg.Gin.Mode = os.Getenv("GIN_MODE")
}

func setAppConfig(cfg *AppConfig) {
	cfg.App.Port = os.Getenv("APP_PORT")
	cfg.App.Name = os.Getenv("APP_NAME")
	cfg.App.Env = os.Getenv("APP_ENV")
}

func setPostgresDbConfig(cfg *AppConfig) error {
	cfg.PostgresDatabase.Username = os.Getenv("POSTGRES_DATABASE_USERNAME")
	cfg.PostgresDatabase.Password = os.Getenv("POSTGRES_DATABASE_PASSWORD")
	cfg.PostgresDatabase.Host = os.Getenv("POSTGRES_DATABASE_HOST")
	cfg.PostgresDatabase.Name = os.Getenv("POSTGRES_DATABASE_NAME")

	port, err := envConvertor("POSTGRES_DATABASE_PORT", func(v string) (uint64, error) {
		return strconv.ParseUint(v, 10, 32)
	})
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.Port = port

	cfg.PostgresDatabase.SslMode = os.Getenv("POSTGRES_DATABASE_SSLMODE")
	cfg.PostgresDatabase.Timezone = os.Getenv("POSTGRES_DATABASE_TIMEZONE")

	maxConn, err := envConvertor("POSTGRES_DATABASE_MAX_OPEN_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.MaxOpenConnections = maxConn

	maxIdle, err := envConvertor("POSTGRES_DATABASE_MAX_IDLE_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.MaxIdleConnections = maxIdle

	connMaxLif, err := envConvertor("POSTGRES_DATABASE_CONN_MAX_LIFETIME", time.ParseDuration)
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.ConnectionMaxLifetime = connMaxLif

	cfg.PostgresDatabase.MigrationPath = os.Getenv("POSTGRES_DATABASE_MIGRATION_PATH")

	return nil
}

func envConvertor[T any](envKey string, converter func(v string) (T, error)) (T, error) {
	value := os.Getenv(envKey)

	result, err := converter(value)
	if err != nil {
		var noop T

		return noop, fmt.Errorf("%s is not a valid value for %s, %w", value, envKey, err)
	}

	return result, nil
}
