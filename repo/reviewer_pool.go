package repo

import (
	"github.com/jackc/pgx"
	"log"
)

func NewReviewerPool(username, database string) *pgx.ConnPool {
	poolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			User:     username,
			Password: "",
			Database: database,
		},
		MaxConnections: 5,
	}
	pool, err := pgx.NewConnPool(poolConfig)
	if err != nil {
		log.Fatal("Unable to config connection pool")
	}
	return pool
}