package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j4kubM/basketball-api/handler"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "m=ri-*GI8RoqE?I1lj8f"
	dbname   = "basketballdb"
)

func New() http.Handler {
	router := gin.Default()

	db, err := DatabaseConnector()
	CheckError(err)

	h := handler.NewBaseHandler(db)

	router.GET("/team", h.GetTeam)
	router.GET("/player", h.GetPlayers)
	router.POST("/player", h.CreatePlayer)

	return router
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
