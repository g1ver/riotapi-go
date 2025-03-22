package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get all champion mastery entries sorted by number of champion points descending.
func (c *Client) GetAllChampionMasteries(puuid string) ([]models.ChampionMastery, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s", c.baseUrl, puuid)

	var cms []models.ChampionMastery
	resp, err := c.Get(url, &cms)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get champion masteries",
			Headers:    resp.Header,
		}
	}

	return cms, nil
}

// Get a champion mastery by puuid and champion ID.
func (c *Client) GetChampionMastery(puuid string, championId int) (*models.ChampionMastery, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%d", c.baseUrl, puuid, championId)

	var cm models.ChampionMastery
	resp, err := c.Get(url, &cm)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get champion mastery",
			Headers:    resp.Header,
		}
	}

	return &cm, nil
}

// Get specified number of top champion mastery entries sorted by number of champion points descending.
func (c *Client) GetTopChampionMasteries(puuid string, count int) ([]models.ChampionMastery, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top?count=%d", c.baseUrl, puuid, count)

	var cms []models.ChampionMastery
	resp, err := c.Get(url, &cms)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get champion masteries",
			Headers:    resp.Header,
		}
	}

	return cms, nil
}

// Get a player's total champion mastery score, which is the sum of individual champion mastery levels.
func (c *Client) GetMasteryScore(puuid string) (int, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/scores/by-puuid/%s", c.baseUrl, puuid)

	var ms int
	resp, err := c.Get(url, &ms)
	if err != nil {
		return -1, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get mastery score",
			Headers:    resp.Header,
		}
	}

	return ms, nil
}
