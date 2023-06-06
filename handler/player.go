package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j4kubM/basketball-api/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "m=ri-*GI8RoqE?I1lj8f"
	dbname   = "basketballdb"
)

func GetPlayers(ctx *gin.Context) {

	db, err := DatabaseConnector()
	if err != nil {
		ctx.Error(err)
	}

	rows, err := db.Query("SELECT * FROM players ")
	if err != nil {
		ctx.Error(err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	var players []model.Player
	fmt.Println(rows)
	for rows.Next() {
		var player model.Player
		if err := rows.Scan(
			&player.ID,
			&player.FirstName,
			&player.LastName,
			&player.Height,
			&player.Number,
			&player.Nationality,
			&player.DateOfBirth); err != nil {
			ctx.Error(err)
		}
		players = append(players, player)
	}
	db.Close()

	ctx.IndentedJSON(http.StatusOK, players)
}

func DatabaseConnector() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf(
		`host=%s 
	port=%d 
	user=%s 
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	fmt.Println("Connected!")
	return db, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
