package repository

import (
	"context"
	"github.com/Sirok47/CarsServer/model"
	"github.com/jackc/pgx"
)

func (r Cars) CreateCar(ctx context.Context, g model.CarParams) error {
	emp := &pgx.QueryExOptions{}
	_, err := r.db.ExecEx(ctx, "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", emp, g.Carbrand, g.Carnumber, g.Cartype, g.Mileage)
	return err
}

func (r Cars) GetCar(ctx context.Context, num int) (*model.CarParams, error) {
	c := &model.CarParams{Carnumber: num}
	emp := &pgx.QueryExOptions{}
	result, err := r.db.QueryEx(ctx, "select * from cars where carnumber = $1", emp, c.Carnumber)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		err = result.Scan(&c.Carbrand, &c.Carnumber, &c.Cartype, &c.Mileage)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (r Cars) UpdateCar(ctx context.Context, c *model.CarParams) error {
	emp := &pgx.QueryExOptions{}
	_, err := r.db.ExecEx(ctx, "update cars set mileage = $1 where carnumber = $2", emp, c.Mileage, c.Carnumber)
	return err
}

func (r Cars) DeleteCar(ctx context.Context, num int) error {
	emp := &pgx.QueryExOptions{}
	_, err := r.db.ExecEx(ctx, "delete from cars where carnumber = $1", emp, num)
	return err
}
