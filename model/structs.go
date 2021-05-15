package model

type Car struct {
	CarBrand  string `json:"CarBrand"`
	CarNumber int    `json:"CarNumber"`
	CarType   string `json:"Type"`
	Mileage   int    `json:"Mileage"`
}

type User struct {
	Nick     string `json:"Nick"`
	Password string `json:"Password"`
}
