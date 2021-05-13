package repository

import (
	"context"
	"github.com/Sirok47/CarsServer/model"
	"github.com/jackc/pgx/v4"
)

type Cars struct {
	db *pgx.Conn
}

func NewCars(db *pgx.Conn) *Cars {
	return &Cars{db: db}
}

func (r Cars) Create(ctx context.Context, g model.CarParams) error {
	_, err := r.db.Exec(ctx, "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", g.CarBrand, g.CarNumber, g.CarType, g.Mileage)
	return err
}

func (r Cars) Get(ctx context.Context, num int) (*model.CarParams, error) {
	c := &model.CarParams{CarNumber: num}
	result, err := r.db.Query(ctx, "select * from cars where carnumber = $1", c.CarNumber)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		err = result.Scan(&c.CarBrand, &c.CarNumber, &c.CarType, &c.Mileage)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (r Cars) Update(ctx context.Context, c *model.CarParams) error {
	_, err := r.db.Exec(ctx, "update cars set mileage = $1 where carnumber = $2", c.Mileage, c.CarNumber)
	return err
}

func (r Cars) Delete(ctx context.Context, num int) error {
	_, err := r.db.Exec(ctx, "delete from cars where carnumber = $1", num)
	return err
}
