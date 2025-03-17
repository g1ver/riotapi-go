package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get League of Legends status for the given platform.
func (c *Client) GetLeagueStatus() (*models.PlatformData, error) {
	url := fmt.Sprintf("%s/lol/status/v4/platform-data", c.baseUrl)

	var pd models.PlatformData
	if err := c.Get(url, &pd); err != nil {
		return nil, fmt.Errorf("failed to get league status: %w", err)
	}

	return &pd, nil
}
