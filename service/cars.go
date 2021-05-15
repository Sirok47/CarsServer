package service

import (
	"context"
	"github.com/Sirok47/CarsServer/model"
	protocol "github.com/Sirok47/CarsServer/protocol"
	"github.com/Sirok47/CarsServer/repository"
	"github.com/jackc/pgx/v4"
)

type Cars struct {
	rps *repository.Cars
}

func NewService(db *pgx.Conn) *Cars {
	return &Cars{repository.NewCars(db)}
}

func (c Cars) SignUp(ctx context.Context, prm *protocol.Userdata) (*protocol.Errmsg, error) {
	err := c.rps.SignUp(ctx, &model.User{Nick: prm.Nick, Password: prm.Password})
	if err != nil {
		return &protocol.Errmsg{Error: err.Error()}, nil
	}
	return nil, nil
}

func (c Cars) LogIn(ctx context.Context, prm *protocol.Userdata) (*protocol.Token, error) {
	token, err := c.rps.LogIn(ctx, &model.User{Nick: prm.Nick, Password: prm.Password})
	if err != nil {
		return &protocol.Token{}, err
	}
	return &protocol.Token{Token: token}, nil
}

func (c Cars) Create(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err := c.rps.Create(ctx, &model.Car{CarBrand: prm.CarBrand, CarNumber: int(prm.CarNumber), CarType: prm.CarType, Mileage: int(prm.Mileage)})
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
	err := c.rps.Update(ctx, &model.Car{CarNumber: int(prm.CarNumber), Mileage: int(prm.Mileage)})
	if err != nil {
		return &protocol.Errmsg{Error: err.Error()}, nil
	}
	return nil, nil
}

func (c Cars) Get(ctx context.Context, prm *protocol.Carparams) (*protocol.Carparams, error) {
	car, err := c.rps.Get(ctx, int(prm.CarNumber))
	if err != nil {
		return &protocol.Carparams{}, err
	}
	return &protocol.Carparams{CarNumber: int32(car.CarNumber), Mileage: int32(car.Mileage), CarType: car.CarType, CarBrand: car.CarBrand}, nil
}
