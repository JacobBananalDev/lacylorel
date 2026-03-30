package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://admin:password@localhost:5432/lacylorel")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn
}