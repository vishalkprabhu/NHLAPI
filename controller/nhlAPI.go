package controller

import (
	"NHLAPI/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const BASE_URL = "https://statsapi.web.nhl.com/api/v1"

func GetTeams() ([]model.Team, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams/", BASE_URL))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response []model.NhlTeamsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	return response.Teams, err

}

func GetTeamsResponse(c echo.Context) error {
	teams, _ := GetTeams()
	return c.JSON(http.StatusOK, teams)
}
