package request

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/j4kubM/basketball-api/model"
)

func ParseCreatePlayerRequest(ctx *gin.Context) (model.PlayerRequest, error) {
	var req model.PlayerRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		return req, fmt.Errorf("failed to parse create player request body: %s", err)
	}

	if req.FirstName == "" {
		return req, fmt.Errorf("players first name cannot be empty")
	}
	fmt.Println(req)
	return req, nil
}
