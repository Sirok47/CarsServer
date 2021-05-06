package service

import (
	"github.com/Sirok47/CarsServer/model"
	"github.com/Sirok47/CarsServer/protocol"
	"context"
)

func (c Cars) CreateCar(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err:=c.rps.CreateCar(ctx,model.CarParams{Carbrand: prm.Carbrand, Carnumber: int(prm.Carnumber), Cartype: prm.Cartype, Mileage: int(prm.Mileage)})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c Cars) DeleteCar(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err:=c.rps.DeleteCar(ctx,int(prm.Carnumber))
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c Cars) UpdateCar(ctx context.Context, prm *protocol.Carparams) (*protocol.Errmsg, error) {
	err:=c.rps.UpdateCar(ctx,&model.CarParams{Carnumber: int(prm.Carnumber), Mileage: int(prm.Mileage)})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c Cars) GetCar(ctx context.Context, prm *protocol.Carparams) (*protocol.Carparams, error) {
	car,err:=c.rps.GetCar(ctx,int(prm.Carnumber))
	if err!=nil{
		return &protocol.Carparams{}, err
	}
	return &protocol.Carparams{Carnumber: int32(car.Carnumber),Mileage: int32(car.Mileage),Cartype: car.Cartype,Carbrand: car.Carbrand}, nil
}