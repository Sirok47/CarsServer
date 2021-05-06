package repository

import "github.com/jackc/pgx"

type Cars struct {
	db *pgx.Conn
}

func NewCars(db *pgx.Conn) *Cars {
	return &Cars{db: db}
}
