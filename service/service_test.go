package service

import (
	"context"
	"errors"
	"github.com/Sirok47/CarsServer/model"
	protocol "github.com/Sirok47/CarsServer/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mocking struct {
	mock.Mock
}

func (m *mocking) SignUp(ctx context.Context, nick string, password string) error {
	args := m.Called(ctx, nick, password)
	return args.Error(0)
}

func (m *mocking) LogIn(ctx context.Context, nick string, password string) (string, error) {
	args := m.Called(ctx, nick, password)
	return args.String(0), args.Error(1)
}

func (m *mocking) Create(ctx context.Context, brand string, num int, car_type string, mileage int) error {
	args := m.Called(ctx, brand, num, car_type, mileage)
	return args.Error(0)
}

func (m *mocking) Update(ctx context.Context, num int, mileage int) error {
	args := m.Called(ctx, num, mileage)
	return args.Error(0)
}

func (m *mocking) Get(ctx context.Context, num int) (*model.Car, error) {
	args := m.Called(ctx, num)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *mocking) Delete(ctx context.Context, num int) error {
	args := m.Called(ctx, num)
	return args.Error(0)
}

func TestSignUp(t *testing.T) {
	testobj := &mocking{}
	testobj.On("SignUp", context.Background(), "test", "test").Return(nil)
	test := Cars{rps: testobj}
	_, err := test.SignUp(context.Background(), &protocol.Userdata{Nick: "test", Password: "test"})
	assert.Equal(t, nil, err)
}

func TestSignUp_Error(t *testing.T) {
	testobj := &mocking{}
	testobj.On("SignUp", context.Background(), "test", "test").Return(errors.New("testerror"))
	test := Cars{rps: testobj}
	msg, err := test.SignUp(context.Background(), &protocol.Userdata{Nick: "test", Password: "test"})
	if msg != nil && err != nil {
		assert.Equal(t, msg.Error, err.Error())
	} else {
		t.Errorf("Got no error")
	}
}

func TestLogIn(t *testing.T) {
	testobj := &mocking{}
	testobj.On("LogIn", context.Background(), "test", "test").Return("testtoken", nil)
	test := Cars{rps: testobj}
	token, err := test.LogIn(context.Background(), &protocol.Userdata{Nick: "test", Password: "test"})
	assert.NotEqual(t, nil, token)
	assert.Equal(t, nil, err)
}

func TestLogIn_Error(t *testing.T) {
	testobj := &mocking{}
	testobj.On("LogIn", context.Background(), "test", "test").Return("", errors.New("testerror"))
	test := Cars{rps: testobj}
	token, err := test.LogIn(context.Background(), &protocol.Userdata{Nick: "test", Password: "test"})
	if token != nil {
		t.Errorf("Expected nil token, got %v", token.Token)
	}
	assert.NotEqual(t, nil, err)
}

func TestCreate(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Create", context.Background(), "test", 1111, "test", 1111).Return(nil)
	test := Cars{rps: testobj}
	_, err := test.Create(context.Background(), &protocol.Carparams{CarNumber: 1111, CarType: "test", CarBrand: "test", Mileage: 1111})
	assert.Equal(t, nil, err)
}

func TestCreate_Error(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Create", context.Background(), "test", 1111, "test", 1111).Return(errors.New("testerror"))
	test := Cars{rps: testobj}
	msg, err := test.Create(context.Background(), &protocol.Carparams{CarNumber: 1111, CarType: "test", CarBrand: "test", Mileage: 1111})
	if msg != nil && err != nil {
		assert.Equal(t, msg.Error, err.Error())
	} else {
		t.Errorf("Got no error")
	}
}

func TestUpdate(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Update", context.Background(), 1111, 1111).Return(nil)
	test := Cars{rps: testobj}
	_, err := test.Update(context.Background(), &protocol.Carparams{CarNumber: 1111, Mileage: 1111})
	assert.Equal(t, nil, err)
}

func TestUpdate_Error(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Update", context.Background(), 1111, 1111).Return(errors.New("testerror"))
	test := Cars{rps: testobj}
	msg, err := test.Update(context.Background(), &protocol.Carparams{CarNumber: 1111, Mileage: 1111})
	if msg != nil && err != nil {
		assert.Equal(t, msg.Error, err.Error())
	} else {
		t.Errorf("Got no error")
	}
}

func TestGet(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Get", context.Background(), 1111).Return(&model.Car{CarNumber: 1111, CarType: "test", CarBrand: "test", Mileage: 1111}, nil)
	test := Cars{rps: testobj}
	car, err := test.Get(context.Background(), &protocol.Carparams{CarNumber: 1111})
	assert.NotEqual(t, nil, car)
	assert.Equal(t, nil, err)
}

func TestGet_Error(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Get", context.Background(), 1111).Return(&model.Car{}, errors.New("testerror"))
	test := Cars{rps: testobj}
	car, err := test.Get(context.Background(), &protocol.Carparams{CarNumber: 1111})
	if car != nil {
		t.Errorf("Expected nil object, got %v", car)
	}
	assert.NotEqual(t, nil, err)
}

func TestDelete(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Delete", context.Background(), 1111).Return(nil)
	test := Cars{rps: testobj}
	_, err := test.Delete(context.Background(), &protocol.Carparams{CarNumber: 1111})
	assert.Equal(t, nil, err)
}

func TestDelete_Error(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Delete", context.Background(), 1111).Return(errors.New("testerror"))
	test := Cars{rps: testobj}
	msg, err := test.Delete(context.Background(), &protocol.Carparams{CarNumber: 1111})
	if msg != nil && err != nil {
		assert.Equal(t, msg.Error, err.Error())
	} else {
		t.Errorf("Got no error")
	}
}
