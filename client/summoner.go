package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get a summoner by account id.
func (c *Client) GetSummonerByAccountID(accountId string) (*models.Summoner, error) {
	url := fmt.Sprintf("%s/lol/summoner/v4/summoners/by-account/%s", c.baseUrl, accountId)

	var summoner models.Summoner
	resp, err := c.Get(url, &summoner)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get summoner by account ID",
			Headers:    resp.Header,
		}
	}

	return &summoner, nil
}

// Get a summoner by PUUID.
func (c *Client) GetSummonerByPUUID(puuid string) (*models.Summoner, error) {
	url := fmt.Sprintf("%s/lol/summoner/v4/summoners/by-puuid/%s", c.baseUrl, puuid)

	var summoner models.Summoner
	resp, err := c.Get(url, &summoner)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get summoner by PUUID",
			Headers:    resp.Header,
		}
	}

	return &summoner, nil
}

// Get a summoner by summoner ID.
func (c *Client) GetSummonerBySummonerID(summonerId string) (*models.Summoner, error) {
	url := fmt.Sprintf("%s/lol/summoner/v4/summoners/%s", c.baseUrl, summonerId)

	var summoner models.Summoner
	resp, err := c.Get(url, &summoner)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get summoner by summoner ID",
			Headers:    resp.Header,
		}
	}

	return &summoner, nil
}
