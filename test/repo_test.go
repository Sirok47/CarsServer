package test

import (
	"context"
	"github.com/Sirok47/CarsServer/model"
	"github.com/Sirok47/CarsServer/repository"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v4"
	"os"
	"testing"
)

var (
	dbconn *pgx.Conn
	r      *repository.Cars
)

func TestMain(m *testing.M) {
	dbconn, _ = pgx.Connect(context.Background(), "postgres://maks:glazirovanniisirok@127.0.0.1:5432/testcars")
	r = repository.NewCars(dbconn)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSignUp(t *testing.T) {
	dbconn.Exec(context.Background(), "insert into users (nick, password) values ($1,$2)", "keklik", "qpwoeirutyM123")
	err := r.SignUp(context.Background(), "tester", "glazirok")
	if err != nil {
		t.Errorf("SignUp() got err %v", err.Error())
	}
	dbconn.Exec(context.Background(), "delete from users")
}

func TestSignUp_Error(t *testing.T) {
	err := r.SignUp(context.Background(), "", "")
	if err == nil {
		t.Errorf("SignUp() got no err witn incorrect nick length")
	}
}

func TestLogIn(t *testing.T) {
	dbconn.Exec(context.Background(), "insert into users (nick, password) values ($1,$2)", "keklik", "qpwoeirutyM123")
	token, err := r.LogIn(context.Background(), "keklik", "qpwoeirutyM123")
	if err != nil {
		t.Errorf("LogIn() got err %v", err.Error())
	}
	if token == "" {
		t.Errorf("LogIn() returned token is empty")
	}
	dbconn.Exec(context.Background(), "delete from users")
}

func TestLogIn_Error(t *testing.T) {
	dbconn.Exec(context.Background(), "insert into users (nick, password) values ($1,$2)", "keklik", "qpwoeirutyM123")
	token, err := r.LogIn(context.Background(), "keklik", "qpwoerutyM123")
	if err == nil {
		t.Errorf("LogIn() no err with incorrect password")
	}
	if token != "" {
		t.Errorf("LogIn() expected empty token, got %v", token)
	}
	dbconn.Exec(context.Background(), "delete from users")
}

func TestCreate(t *testing.T) {
	err := r.Create(context.Background(), "test", 5000, "test", 1000)
	if err != nil {
		t.Errorf("Create() got err %v", err.Error())
	}
	dbconn.Exec(context.Background(), "delete from cars")
}

func TestCreate_Error(t *testing.T) {
	err := r.Create(context.Background(), "test", 470, "test", 1000)
	if err == nil {
		t.Errorf("Create() no err with incorrect number")
	}
	dbconn.Exec(context.Background(), "delete from cars")
}

func TestUpdate(t *testing.T) {
	dbconn.Exec(context.Background(), "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", "brand", 1331, "type", 1111)
	err := r.Update(context.Background(), 1331, 1112)
	if err != nil {
		t.Errorf("Update() got err %v", err.Error())
	}
	dbconn.Exec(context.Background(), "delete from cars")
}

func TestUpdate_Error(t *testing.T) {
	err := r.Update(context.Background(), 0, 1111)
	if err == nil {
		t.Errorf("Update() got no err with nonexisting number")
	}
}

func TestGet(t *testing.T) {
	dbconn.Exec(context.Background(), "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", "brand", 1331, "type", 1111)
	car, err := r.Get(context.Background(), 1331)
	if err != nil {
		t.Errorf("Get() got err %v", err.Error())
	}
	if car == nil {
		t.Errorf("Get() returned value is nil")
		return
	}
	if !cmp.Equal(car, &model.Car{CarBrand: "brand", CarType: "type", CarNumber: 1331, Mileage: 1111}) {
		t.Errorf("Get(): incorrect parsing")
	}
	dbconn.Exec(context.Background(), "delete from cars")
}

func TestGet_Error(t *testing.T) {
	car, err := r.Get(context.Background(), 0)
	if err == nil {
		t.Errorf("got no err getting nonexisting car")
	}
	if car != nil {
		t.Errorf("Got non-nil car: %v", car)
	}
}

func TestDelete(t *testing.T) {
	dbconn.Exec(context.Background(), "insert into cars (carbrand,carnumber,type,mileage) values ($1,$2,$3,$4)", "brand", 1331, "type", 1111)
	err := r.Delete(context.Background(), 1331)
	if err != nil {
		t.Errorf("Delete() got err %v", err.Error())
	}
	dbconn.Exec(context.Background(), "delete from cars")
}

func TestDelete_Error(t *testing.T) {
	err := r.Delete(context.Background(), 0)
	if err == nil {
		t.Errorf("Delete() got no err deleting nonexisting car")
	}
}
