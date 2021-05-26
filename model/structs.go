package model

type Car struct {
	CarBrand  string `json:"CarBrand" validate:"min=3,max=10"`
	CarNumber int    `json:"CarNumber" validate:"required,min=1000,max=9999"`
	CarType   string `json:"Type" validate:"min=3,max=10"`
	Mileage   int    `json:"Mileage" validate:"min=0,max=1000000"`
}

type User struct {
	Nick     string `json:"Nick"  validate:"required,min=3,max=10"`
	Password string `json:"Password"  validate:"required,min=5,max=20"`
}
