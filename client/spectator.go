package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get current game information for the given puuid.
func (c *Client) GetActiveGameInfoByPUUID(puuid string) (*models.CurrentGameInfo, error) {
	u := fmt.Sprintf("%s/lol/spectator/v5/active-games/by-summoner/%s", c.baseUrl, puuid)

	var gameInfo models.CurrentGameInfo
	resp, err := c.Get(u, &gameInfo)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get active game info by puuid",
			Headers:    resp.Header,
		}
	}

	return &gameInfo, nil
}

// Get a list of featured games.
func (c *Client) GetFeaturedGames() (*models.FeaturedGames, error) {
	u := fmt.Sprintf("%s/lol/spectator/v5/featured-games", c.baseUrl)

	var featuredGames models.FeaturedGames
	resp, err := c.Get(u, &featuredGames)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get featured games",
			Headers:    resp.Header,
		}
	}

	return &featuredGames, nil
}
