package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/j4kubM/basketball-api/handler/request"
	"github.com/j4kubM/basketball-api/model"
)

func (h *BaseHandler) GetPlayers(ctx *gin.Context) {

	rows, err := h.db.Query("SELECT * FROM players ")
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

	ctx.IndentedJSON(http.StatusOK, players)
}

func GetPlayersByTeamQuery(int) string {
	query := `SELECT
			*
		FROM players
		WHERE teamId = 1;`

	return query
}

type query string

func createPlayerQuery(player model.PlayerRequest) query {
	return query(fmt.Sprintf(`INSERT INTO players (firstName, lastName, height, number, nationality, teamId, dateOfBirth)
	VALUES ('%s', '%s', '%d', '%d', '%s', 1, '%s');`, player.FirstName, player.LastName, player.Height, player.Number, player.Nationality, &player.DateOfBirth))
}

func (h *BaseHandler) CreatePlayer(ctx *gin.Context) {

	req, err := request.ParseCreatePlayerRequest(ctx)
	if err != nil {
		ctx.Error(err)
		http.Error(ctx.Writer, err.Error(), 401)
	}

	q := createPlayerQuery(req)
	fmt.Println(q)
	h.db.Query(string(q))
}
