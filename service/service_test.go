package service

import (
	"context"
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

func TestCreate(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Create", context.Background(), "test", 1111, "test", 1111).Return(nil)
	test := Cars{rps: testobj}
	_, err := test.Create(context.Background(), &protocol.Carparams{})
	assert.Equal(t, nil, err)
}
