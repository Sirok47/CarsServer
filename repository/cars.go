package repository

import (
	"context"
	"errors"
	"github.com/Sirok47/CarsServer/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"time"
)

type Cars struct {
	db *pgx.Conn
}

func NewCars(db *pgx.Conn) *Cars {
	return &Cars{db: db}
}

type CarsInt interface {
	Create(ctx context.Context, brand string, num int, car_type string, mileage int) error
}

func (r Cars) SignUp(ctx context.Context, nick string, password string) error {
	if nick == "" || len(nick) > 20 {
		return errors.New("incorrect nick length")
	}
	_, err := r.db.Exec(ctx, "insert into users (nick,password) values ($1,$2)", nick, password)
	return err
}

func (r Cars) LogIn(ctx context.Context, nick string, password string) (string, error) {
	trueuser := &model.User{}
	result, err := r.db.Query(ctx, "select * from users where nick = $1", nick)
	if err != nil {
		return "", err
	}
	defer result.Close()
	for result.Next() {
		err = result.Scan(&trueuser.Nick, &trueuser.Password)
		if err != nil {
			return "", err
		}
	}
	if password == trueuser.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour).Unix()
		t, err := token.SignedString([]byte("sirok"))
		if err != nil {
			return "", err
		}
		return t, nil
	}
	return "", echo.ErrUnauthorized
}

func (r Cars) Create(ctx context.Context, brand string, num int, car_type string, mileage int) error {
	if num < 1000 || num > 9999 {
		return errors.New("wrong CarNumber")
	}
	_, err := r.db.Exec(ctx, "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", brand, num, car_type, mileage)
	return err
}

func (r Cars) Get(ctx context.Context, num int) (*model.Car, error) {
	c := &model.Car{CarNumber: num}
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
	if c.CarNumber == 0 {
		return nil, errors.New("nonexistent car")
	}
	return c, nil
}

func (r Cars) Update(ctx context.Context, num int, mileage int) error {
	res, err := r.db.Exec(ctx, "update cars set mileage = $1 where carnumber = $2", mileage, num)
	if res == nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("nonexistent car")
	}
	return err
}

func (r Cars) Delete(ctx context.Context, num int) error {
	res, err := r.db.Exec(ctx, "delete from cars where carnumber = $1", num)
	if res == nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("nonexistent car")
	}
	return err
}
