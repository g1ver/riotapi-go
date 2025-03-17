package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get current game information for the given puuid.
func (c *Client) GetActiveGameInfoByPUUID(puuid string) (*models.CurrentGameInfo, error) {
	u := fmt.Sprintf("%s/lol/spectator/v5/active-games/by-summoner/%s", c.baseUrl, puuid)

	var gameInfo models.CurrentGameInfo
	if err := c.Get(u, &gameInfo); err != nil {
		return nil, err
	}

	return &gameInfo, nil
}

// Get a list of featured games.
func (c *Client) GetFeaturedGames() (*models.FeaturedGames, error) {
	u := fmt.Sprintf("%s/lol/spectator/v5/featured-games", c.baseUrl)

	var featuredGames models.FeaturedGames
	if err := c.Get(u, &featuredGames); err != nil {
		return nil, err
	}

	return &featuredGames, nil
}
