package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j4kubM/basketball-api/mock"
)

func GetPlayers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, mock.Player1)
}
