package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

var validRegions = map[string]bool{
	"americas": true,
	"europe":   true,
	"asia":     true,
	"esports":  true,
}

// Get account by puuid
func (c *Client) GetAccountByPUUID(puuid string, region string) (*models.Account, error) {
	if !validRegions[region] {
		region = "americas"
	}

	url := fmt.Sprintf("https://%s.api.riotgames.com/riot/account/v1/accounts/by-puuid/%s", region, puuid)

	var account models.Account
	if err := c.Get(url, &account); err != nil {
		return nil, fmt.Errorf("failed to get summoner by name: %w", err)
	}

	return &account, nil
}

// Get account by riot id
func (c *Client) GetAccountByRiotID(gameName string, tagLine string, region string) (*models.Account, error) {
	if !validRegions[region] {
		region = "americas"
	}

	url := fmt.Sprintf("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", region, gameName, tagLine)

	var account models.Account
	if err := c.Get(url, &account); err != nil {
		return nil, fmt.Errorf("failed to get summoner by riot id: %w", err)
	}

	return &account, nil
}
