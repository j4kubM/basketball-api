package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j4kubM/basketball-api/handler"
	_ "github.com/lib/pq"
)

func New() http.Handler {
	router := gin.Default()

	router.GET("/team", handler.GetTeam)
	router.GET("/player", handler.GetPlayers)

	return router
}
