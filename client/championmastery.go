package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get all champion mastery entries sorted by number of champion points descending.
func (c *Client) GetAllChampionMasteries(puuid string) ([]models.ChampionMastery, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s", c.baseUrl, puuid)

	var cms []models.ChampionMastery
	if err := c.Get(url, &cms); err != nil {
		return nil, fmt.Errorf("failed to get champion masteries: %w", err)
	}

	return cms, nil
}

// Get a champion mastery by puuid and champion ID.
func (c *Client) GetChampionMastery(puuid string, championId int) (*models.ChampionMastery, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%d", c.baseUrl, puuid, championId)

	var cm models.ChampionMastery
	if err := c.Get(url, &cm); err != nil {
		return nil, fmt.Errorf("failed to get champion mastery: %w", err)
	}

	return &cm, nil
}

// Get specified number of top champion mastery entries sorted by number of champion points descending.
func (c *Client) GetTopChampionMasteries(puuid string, count int) ([]models.ChampionMastery, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top?=%d", c.baseUrl, puuid, count)

	var cms []models.ChampionMastery
	if err := c.Get(url, &cms); err != nil {
		return nil, fmt.Errorf("failed to get champion masteries: %w", err)
	}

	return cms, nil
}

// Get a player's total champion mastery score, which is the sum of individual champion mastery levels.
func (c *Client) GetMasteryScore(puuid string) (int, error) {
	url := fmt.Sprintf("%s/lol/champion-mastery/v4/scores/by-puuid/%s", c.baseUrl, puuid)

	var ms int
	if err := c.Get(url, &ms); err != nil {
		return -1, fmt.Errorf("failed to get champion masteries: %w", err)
	}

	return ms, nil
}
