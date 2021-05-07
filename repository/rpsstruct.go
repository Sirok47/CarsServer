package repository

import "github.com/jackc/pgx/v4"

type Cars struct {
	db *pgx.Conn
}

func NewCars(db *pgx.Conn) *Cars {
	return &Cars{db: db}
}
