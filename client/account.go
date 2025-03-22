package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get account by puuid
func (c *Client) GetAccountByPUUID(puuid string) (*models.Account, error) {
	url := fmt.Sprintf("%s/riot/account/v1/accounts/by-puuid/%s", c.routingServerUrl, puuid)

	var account models.Account
	resp, err := c.Get(url, &account)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get account by PUUID",
			Headers:    resp.Header,
		}
	}
	return &account, nil
}

// Get account by riot id
func (c *Client) GetAccountByRiotID(gameName string, tagLine string) (*models.Account, error) {
	url := fmt.Sprintf("%s/riot/account/v1/accounts/by-riot-id/%s/%s", c.routingServerUrl, gameName, tagLine)

	var account models.Account
	resp, err := c.Get(url, &account)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get account by Riot ID",
			Headers:    resp.Header,
		}
	}

	return &account, nil
}
