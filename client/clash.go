package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get players by puuid.
func (c *Client) GetClashPlayersByPUUID(puuid string) ([]models.Player, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/players/by-puuid/%s", c.baseUrl, puuid)

	var ps []models.Player
	if err := c.Get(url, &ps); err != nil {
		return nil, fmt.Errorf("failed to get player by puuid: %w", err)
	}

	return ps, nil
}

// Get team by ID.
func (c *Client) GetClashTeamByID(teamId string) (*models.Team, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/teams/%s", c.baseUrl, teamId)

	var t models.Team
	if err := c.Get(url, &t); err != nil {
		return nil, fmt.Errorf("failed to get team by id: %w", err)
	}

	return &t, nil
}

// Get all active or upcoming tournaments
func (c *Client) GetClashTournaments() ([]models.Tournament, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/tournaments", c.baseUrl)

	var ts []models.Tournament
	if err := c.Get(url, &ts); err != nil {
		return nil, fmt.Errorf("failed to get team by id: %w", err)
	}

	return ts, nil
}

// Get tournament by team ID
func (c *Client) GetClashTournamentByTeamID(teamId string) (*models.Tournament, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/tournaments/by-team/%s", c.baseUrl, teamId)

	var t models.Tournament
	if err := c.Get(url, t); err != nil {
		return nil, fmt.Errorf("failed to get team by id: %w", err)
	}

	return &t, nil
}

// Get tournament by ID
func (c *Client) GetClashTournamentByID(tournamentId string) (*models.Tournament, error) {
	url := fmt.Sprintf("%s/lol/clash/v1/tournaments/%s", c.baseUrl, tournamentId)

	var t models.Tournament
	if err := c.Get(url, t); err != nil {
		return nil, fmt.Errorf("failed to get team by id: %w", err)
	}

	return &t, nil
}
