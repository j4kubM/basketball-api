package mock

import (
	"time"

	"github.com/j4kubM/basketball-api/model"
)

// var Player1 = model.Player{
// 	Height:      201,
// 	Number:      22,
// 	FirstName:   "Jimmy",
// 	LastName:    "Butler",
// 	DateOfBirth: model.Date{time.Date(1989, 9, 14, 0, 0, 0, 0, time.UTC)},
// 	Nationality: "USA",
// }

// var Player2 = model.Player{
// 	Height:      217,
// 	Number:      13,
// 	FirstName:   "Bam",
// 	LastName:    "Adebayo",
// 	DateOfBirth: model.Date{time.Date(1997, 6, 18, 0, 0, 0, 0, time.UTC)},
// 	Nationality: "USA",
// }

var Coach1 = model.Coach{
	FirstName:   "Erik",
	LastName:    "Spoelstra",
	DateOfBirth: model.Date(time.Date(1970, 11, 1, 0, 0, 0, 0, time.UTC)),
	Nationality: "USA",
}

var Team1 = model.Team{
	Name: "Miami Heat",
	Address: model.Address{
		StreetName:  "Avenue",
		HouseNumber: 2,
		City:        "Miami",
		Country:     "USA",
	},
	Coach:   Coach1,
	Country: "USA",
	Players: []model.Player{
		// Player1,
		// Player2,
	},
}
