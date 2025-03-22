package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get all the league entries
// Technically, these are all unique players and Riot returns a Set
// I'm just going to use a slice for now
func (c *Client) GetAllLeagueExpEntries(queue string, tier string, division string, page int) ([]models.LeagueEntry, error) {
	url := fmt.Sprintf("%s/lol/league-exp/v4/entries/%s/%s/%s?page=%d", c.baseUrl, queue, tier, division, page)

	var les []models.LeagueEntry
	resp, err := c.Get(url, &les)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get league entries",
			Headers:    resp.Header,
		}
	}

	return les, nil
}
