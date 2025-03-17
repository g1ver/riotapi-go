package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Returns champion rotations, including free-to-play and low-level free-to-play rotations
func (c *Client) GetChampionRotations() (*models.ChampionInfo, error) {
	url := fmt.Sprintf("%s/lol/platform/v3/champion-rotations", c.baseUrl)

	var ci models.ChampionInfo
	if err := c.Get(url, &ci); err != nil {
		return nil, fmt.Errorf("failed to get champion rotations: %w", err)
	}

	return &ci, nil
}
