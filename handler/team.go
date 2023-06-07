package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j4kubM/basketball-api/mock"
)

func (h *BaseHandler) GetTeam(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, mock.Team1)
}
