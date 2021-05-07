package repository

import (
	"context"
	"github.com/Sirok47/CarsServer/model"
)

func (r Cars) CreateCar(ctx context.Context, g model.CarParams) error {
	_, err := r.db.Exec(ctx, "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", g.Carbrand, g.Carnumber, g.Cartype, g.Mileage)
	return err
}

func (r Cars) GetCar(ctx context.Context, num int) (*model.CarParams, error) {
	c := &model.CarParams{Carnumber: num}
	result, err := r.db.Query(ctx, "select * from cars where carnumber = $1", c.Carnumber)
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
	_, err := r.db.Exec(ctx, "update cars set mileage = $1 where carnumber = $2", c.Mileage, c.Carnumber)
	return err
}

func (r Cars) DeleteCar(ctx context.Context, num int) error {
	_, err := r.db.Exec(ctx, "delete from cars where carnumber = $1", num)
	return err
}
