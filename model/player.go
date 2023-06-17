package model

type Player struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth Date   `json:"dateOfBirth"`
	Nationality string `json:"nationality"`
	Height      int    `json:"height"`
	Number      int    `json:"number"`
}

type PlayerRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth Date   `json:"dateOfBirth"`
	Nationality string `json:"nationality"`
	Height      int    `json:"height"`
	Number      int    `json:"number"`
}

type Coach struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth Date   `json:"dateOfBirth"`
	Nationality string `json:"nationality"`
}
