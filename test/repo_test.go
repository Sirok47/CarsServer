package test

import (
	"context"
	"github.com/Sirok47/CarsServer/model"
	"github.com/Sirok47/CarsServer/repository"
	"github.com/jackc/pgx/v4"
	"testing"
)

func TestRepository(t *testing.T) {
	dbconn, _ := pgx.Connect(context.Background(), "postgres://maks:glazirovanniisirok@127.0.0.1:5432/cars")
	r := repository.NewCars(dbconn)
	err := r.SignUp(context.Background(), &model.User{Nick: "keklik_tester", Password: "glazirok"})
	if err != nil {
		t.Errorf("SignUp() got err %v", err.Error())
	}
	_, err = dbconn.Exec(context.Background(), "delete from users where nick = keklik_tester")
	if err != nil {
		t.Errorf("DB issue while deleting")
	}
	_, err = r.LogIn(context.Background(), &model.User{Nick: "keklik", Password: "qpwoeirutyM123"})
	if err != nil {
		t.Errorf("LogIn() got err %v", err.Error())
	}
	err = r.Create(context.Background(), &model.Car{CarBrand: "test", CarNumber: 1331, CarType: "test", Mileage: 1000})
	if err != nil {
		t.Errorf("Create() got err %v", err.Error())
	}
	err = r.Update(context.Background(), &model.Car{CarNumber: 1331, Mileage: 1111})
	if err != nil {
		t.Errorf("Update() got err %v", err.Error())
	}
	_, err = r.Get(context.Background(), 1331)
	if err != nil {
		t.Errorf("Get() got err %v", err.Error())
	}
	err = r.Delete(context.Background(), 1331)
	if err != nil {
		t.Errorf("Delete() got err %v", err.Error())
	}
}
