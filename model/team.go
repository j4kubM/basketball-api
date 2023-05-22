package model

type Team struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	League  int    `json:"league"`
	Coach   Coach
	Address Address
	Players []Player
}
