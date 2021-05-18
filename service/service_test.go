package service

import (
	"context"
	protocol "github.com/Sirok47/CarsServer/protocol"
	"github.com/Sirok47/CarsServer/repository"
	"github.com/jackc/pgx/v4"

	//"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mocking struct {
	mock.Mock
}

func (m *mocking) Create(ctx context.Context, brand string, num int, car_type string, mileage int) error {
	args := m.Called(ctx, brand, num, car_type, mileage)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	testobj := &mocking{}
	testobj.On("Create", context.Background(), "test", 1111, "test", 1111).Return(nil)
	test := Cars{rps: testobj}
	_, err := test.Create(context.Background(), &protocol.Carparams{})
	assert.Equal(t, nil, err)
}
