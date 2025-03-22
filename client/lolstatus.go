package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get League of Legends status for the given platform.
func (c *Client) GetLeagueStatus() (*models.PlatformData, error) {
	url := fmt.Sprintf("%s/lol/status/v4/platform-data", c.baseUrl)

	var pd models.PlatformData
	resp, err := c.Get(url, &pd)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get league status",
			Headers:    resp.Header,
		}
	}

	return &pd, nil
}
