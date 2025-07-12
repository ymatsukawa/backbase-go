package plugin

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ymatsukawa/backbase-go/pkg/env"
)

type Database interface {
	Launch() (*pgx.Conn, error)
}

type database struct{}

func NewDatabase() Database {
	return &database{}
}

func (d *database) Launch() (*pgx.Conn, error) {
	ctx := context.Background()

	// Build database connection string from environment variables
	host := env.GetEnv("DB_HOST", "pdb")
	port := env.GetEnv("DB_PORT", "5432")
	user := env.GetEnv("DB_USER", "root")
	password := env.GetEnv("DB_PASSWORD", "password")
	dbname := env.GetEnv("DB_NAME", "root")
	sslmode := env.GetEnv("DB_SSL_MODE", "disable")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return conn, nil
}
