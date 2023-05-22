package model

import "time"

type Player struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Nationality string    `json:"nationality"`
	Height      int       `json:"height"`
	Number      int       `json:"number"`
}

type Coach struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Nationality string    `json:"nationality"`
}
