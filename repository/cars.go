package repository

import (
	"context"
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

func (r Cars) SignUp(ctx context.Context, g model.UserParams) error {
	_, err := r.db.Exec(ctx, "insert into users (nick,password) values ($1,$2)", g.Nick, g.Password)
	return err
}

func (r Cars) LogIn(ctx context.Context, user model.UserParams) (string, error) {
	trueuser := model.UserParams{}
	result, err := r.db.Query(ctx, "select * from users where nick = $1", user.Nick)
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
	if user == trueuser {
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
