package service

import (
	"context"
	protocol "github.com/Sirok47/CarsServer/protocol"
	"github.com/Sirok47/CarsServer/repository"
	"github.com/jackc/pgx/v4"
)

type Cars struct {
	rps repository.CarsInt
}

func NewService(db *pgx.Conn) *Cars {
	return &Cars{rps: repository.NewCars(db)}
}

func (c Cars) SignUp(ctx context.Context, prm *protocol.Userdata) (*protocol.Errmsg, error) {
	err := c.rps.SignUp(ctx, prm.Nick, prm.Password)
	if err != nil {
		return &protocol.Errmsg{Error: err.Error()}, nil
	}
	return nil, nil
}

func (c Cars) LogIn(ctx context.Context, prm *protocol.Userdata) (*protocol.Token, error) {
	token, err := c.rps.LogIn(ctx, prm.Nick, prm.Password)
	if err != nil || token == "" {
		return nil, err
	}
	return &protocol.Token{Token: token}, nil
}

func (c Cars) Create(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err := c.rps.Create(ctx, prm.CarBrand, int(prm.CarNumber), prm.CarType, int(prm.Mileage))
	if err != nil {
		return &protocol.Errmsg{Error: err.Error()}, nil
	}
	return nil, nil
}

func (c Cars) Delete(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err := c.rps.Delete(ctx, int(prm.CarNumber))
	if err != nil {
		return &protocol.Errmsg{Error: err.Error()}, nil
	}
	return nil, nil
}

func (c Cars) Update(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err := c.rps.Update(ctx, int(prm.CarNumber), int(prm.Mileage))
	if err != nil {
		return &protocol.Errmsg{Error: err.Error()}, nil
	}
	return nil, nil
}

func (c Cars) Get(ctx context.Context, prm *protocol.Carparams) (*protocol.Carparams, error) {
	car, err := c.rps.Get(ctx, int(prm.CarNumber))
	if err != nil || car == nil {
		return nil, err
	}
	return &protocol.Carparams{CarNumber: int32(car.CarNumber), Mileage: int32(car.Mileage), CarType: car.CarType, CarBrand: car.CarBrand}, nil
}
