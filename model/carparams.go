package model

type CarParams struct {
	CarBrand  string `json:"CarBrand"`
	CarNumber int    `json:"CarNumber"`
	CarType   string `json:"Type"`
	Mileage   int    `json:"Mileage"`
}
