package client

import (
	"fmt"

	"github.com/g1ver/riotapi-go/models"
)

// Get the challenger league for a given queue.
func (c *Client) GetChallengerLeagueByQueue(queue string) (*models.LeagueList, error) {
	url := fmt.Sprintf("%s/lol/league/v4/challengerleagues/by-queue/%s", c.baseUrl, queue)

	var ll models.LeagueList
	resp, err := c.Get(url, &ll)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get challenger league by queue",
			Headers:    resp.Header,
		}
	}

	return &ll, nil
}

// Get league entries in all queues for a given puuid
func (c *Client) GetAllLeagueEntriesByPUUID(puuid string) ([]models.LeagueEntry, error) {
	url := fmt.Sprintf("%s/lol/league/v4/entries/by-puuid/%s", c.baseUrl, puuid)

	var le []models.LeagueEntry
	resp, err := c.Get(url, &le)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get all league entries by puuid",
			Headers:    resp.Header,
		}
	}

	return le, nil
}

// Get league entries in all queues for a given summoner ID.
func (c *Client) GetAllLeagueEntriesBySummonerID(summonerID string) ([]models.LeagueEntry, error) {
	url := fmt.Sprintf("%s/lol/league/v4/entries/by-summoner/%s", c.baseUrl, summonerID)

	var le []models.LeagueEntry
	resp, err := c.Get(url, &le)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get all league entries by summoner id",
			Headers:    resp.Header,
		}
	}

	return le, nil
}

// Get all the league entries.
func (c *Client) GetAllLeagueEntries(queue string, tier string, division string, page int) ([]models.LeagueEntry, error) {
	url := fmt.Sprintf("%s/lol/league/v4/entries/%s/%s/%s?page=%d", c.baseUrl, queue, tier, division, page)

	var les []models.LeagueEntry
	resp, err := c.Get(url, &les)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get all league entries",
			Headers:    resp.Header,
		}
	}

	return les, nil
}

// Get the grandmaster league of a specific queue.
func (c *Client) GetGrandmasterLeagueByQueue(queue string) (*models.LeagueList, error) {
	url := fmt.Sprintf("%s/lol/league/v4/grandmasterleagues/by-queue/%s", c.baseUrl, queue)

	var ll models.LeagueList
	resp, err := c.Get(url, &ll)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get grandmaster league by queue",
			Headers:    resp.Header,
		}
	}

	return &ll, nil
}

// Get league with given ID, including inactive entries.
func (c *Client) GetLeagueByID(leagueID string) (*models.LeagueList, error) {
	url := fmt.Sprintf("%s/lol/league/v4/leagues/%s", c.baseUrl, leagueID)

	var ll models.LeagueList
	resp, err := c.Get(url, &ll)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get league by ID",
			Headers:    resp.Header,
		}
	}

	return &ll, nil
}

// Get the master league for given queue.
func (c *Client) GetMasterLeagueByQueue(queue string) (*models.LeagueList, error) {
	url := fmt.Sprintf("%s/lol/league/v4/masterleagues/by-queue/%s", c.baseUrl, queue)

	var ll models.LeagueList
	resp, err := c.Get(url, &ll)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get master league by queue",
			Headers:    resp.Header,
		}
	}

	return &ll, nil
}
