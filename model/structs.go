package model

type Car struct {
	CarBrand  string `json:"CarBrand" validate:"omitempty,min=3,max=10"`
	CarNumber int    `json:"CarNumber" validate:"required,numeric,len=4"`
	CarType   string `json:"Type" validate:"omitempty,min=3,max=10"`
	Mileage   int    `json:"Mileage" validate:"omitempty,numeric,min=0,max=1000000"`
}

type User struct {
	Nick     string `json:"Nick"  validate:"required,min=3,max=10"`
	Password string `json:"Password"  validate:"required,min=5,max=20"`
}
