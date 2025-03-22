package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get players by puuid.
func (c *Client) GetClashPlayersByPUUID(puuid string) ([]models.Player, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/players/by-puuid/%s", c.baseUrl, puuid)

	var ps []models.Player
	resp, err := c.Get(url, &ps)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get player by puuid",
			Headers:    resp.Header,
		}
	}

	return ps, nil
}

// Get team by ID.
func (c *Client) GetClashTeamByID(teamId string) (*models.ClashTeam, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/teams/%s", c.baseUrl, teamId)

	var t models.ClashTeam
	resp, err := c.Get(url, &t)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get team by id",
			Headers:    resp.Header,
		}
	}

	return &t, nil
}

// Get all active or upcoming tournaments
func (c *Client) GetClashTournaments() ([]models.Tournament, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/tournaments", c.baseUrl)

	var ts []models.Tournament
	resp, err := c.Get(url, &ts)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get tournaments",
			Headers:    resp.Header,
		}
	}

	return ts, nil
}

// Get tournament by team ID
func (c *Client) GetClashTournamentByTeamID(teamId string) (*models.Tournament, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/tournaments/by-team/%s", c.baseUrl, teamId)

	var t models.Tournament
	resp, err := c.Get(url, &t)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get tournament by team id",
			Headers:    resp.Header,
		}
	}

	return &t, nil
}

// Get tournament by ID
func (c *Client) GetClashTournamentByID(tournamentId string) (*models.Tournament, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/tournaments/%s", c.baseUrl, tournamentId)

	var t models.Tournament
	resp, err := c.Get(url, &t)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get tournament by id",
			Headers:    resp.Header,
		}
	}

	return &t, nil
}
