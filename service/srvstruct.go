package service

import (
	"github.com/Sirok47/CarsServer/repository"
	"github.com/jackc/pgx/v4"
)

type Cars struct {
	rps *repository.Cars
}

func NewService(db *pgx.Conn) *Cars {
	return &Cars{repository.NewCars(db)}
}
