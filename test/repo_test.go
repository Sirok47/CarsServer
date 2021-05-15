package test

import (
	"context"
	"github.com/Sirok47/CarsServer/repository"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"os"
	"testing"
)

var (
	dbconn *pgx.Conn
	r      *repository.Cars
	res    pgconn.CommandTag
)

func TestMain(m *testing.M) {
	dbconn, _ = pgx.Connect(context.Background(), "postgres://maks:glazirovanniisirok@127.0.0.1:5432/cars")
	r = repository.NewCars(dbconn)
	os.Exit(m.Run())
}

func TestSignUp(t *testing.T) {
	err := r.SignUp(context.Background(), "tester", "glazirok")
	if err != nil {
		t.Errorf("SignUp() got err %v", err.Error())
	} else {
		res, err = dbconn.Exec(context.Background(), "delete from users where nick = $1", "tester")
		if err != nil || res.RowsAffected() == 0 {
			t.Errorf("DB issue while deleting")
		}
	}
	err = r.SignUp(context.Background(), "", "")
	if err == nil {
		t.Errorf("SignUp() got no err witn incorrect nick length")
	}
}

func TestLogIn(t *testing.T) {
	token, err := r.LogIn(context.Background(), "keklik", "qpwoeirutyM123")
	if err != nil {
		t.Errorf("LogIn() got err %v", err.Error())
	}
	if token == "" {
		t.Errorf("LogIn() returned token is empty")
	}
	token, err = r.LogIn(context.Background(), "keklik", "qpwoerutyM123")
	if err == nil {
		t.Errorf("LogIn() no err with incorrect password")
	}
	if token != "" {
		t.Errorf("LogIn() expected empty token, got %v", token)
	}
}

func TestCreate(t *testing.T) {
	err := r.Create(context.Background(), "test", 1331, "test", 1000)
	if err != nil {
		t.Errorf("Create() got err %v", err.Error())
	}
	err = r.Create(context.Background(), "test", 470, "test", 1000)
	if err == nil {
		t.Errorf("Create() no err with incorrect number")
	}
}

func TestUpdate(t *testing.T) {
	err := r.Update(context.Background(), 1331, 1111)
	if err != nil {
		t.Errorf("Update() got err %v", err.Error())
	}
	err = r.Update(context.Background(), 0, 1111)
	if err == nil {
		t.Errorf("Update() got no err with nonexisting number")
	}
}

func TestGet(t *testing.T) {
	car, err := r.Get(context.Background(), 1331)
	if err != nil {
		t.Errorf("Get() got err %v", err.Error())
	}
	if car == nil {
		t.Errorf("Get() returned value is nil")
	}
	car, err = r.Get(context.Background(), 0)
	if err == nil {
		t.Errorf("got no err getting nonexisting car")
	}
	if car != nil {
		t.Errorf("Got non-nil car: %v", car)
	}
}

func TestDelete(t *testing.T) {
	err := r.Delete(context.Background(), 1331)
	if err != nil {
		t.Errorf("Delete() got err %v", err.Error())
	}
	err = r.Delete(context.Background(), 0)
	if err == nil {
		t.Errorf("Delete() got no err deleting nonexisting car")
	}
}
