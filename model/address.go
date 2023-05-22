package model

type Address struct {
	StreetName           string `json:"streetName"`
	HouseNumber          int    `json:"houseNumber"`
	HouseNumberExtension string `json:"houseNumberExtension,omitempty"`
	City                 string `json:"city"`
	Country              string `json:"country"`
}
