package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Returns champion rotations, including free-to-play and low-level free-to-play rotations
func (c *Client) GetChampionRotations() (*models.ChampionInfo, error) {
	url := fmt.Sprintf("%s/lol/platform/v3/champion-rotations", c.baseUrl)

	var ci models.ChampionInfo
	resp, err := c.Get(url, &ci)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get champion rotations",
			Headers:    resp.Header,
		}
	}

	return &ci, nil
}
