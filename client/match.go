package client

import (
	"fmt"
	"net/url"

	"github.com/g1ver/riotapi-go/models"
)

type GetMatchIDsByPUUIDParams struct {
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
	Queue     int    `url:"queue,omitempty"`
	Type      string `url:"type,omitempty"`
	Start     int    `url:"start,omitempty"`
	Count     int    `url:"count,omitempty"`
}

func (c *Client) GetMatchIDsByPUUID(puuid string, params models.GetMatchIDsByPUUIDParams) ([]string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/lol/match/v5/matches/by-puuid/%s/ids", c.routingServerUrl, puuid))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	if params.StartTime > 0 {
		q.Add("startTime", fmt.Sprintf("%d", params.StartTime))
	}
	if params.EndTime > 0 {
		q.Add("endTime", fmt.Sprintf("%d", params.EndTime))
	}
	if params.Queue > 0 {
		q.Add("queue", fmt.Sprintf("%d", params.Queue))
	}
	if params.Type != "" {
		q.Add("type", params.Type)
	}
	if params.Start > 0 {
		q.Add("start", fmt.Sprintf("%d", params.Start))
	}
	if params.Count > 0 {
		q.Add("count", fmt.Sprintf("%d", params.Count))
	}
	u.RawQuery = q.Encode()

	var matchIDs []string
	resp, err := c.Get(u.String(), &matchIDs)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get match IDs",
			Headers:    resp.Header,
		}
	}

	return matchIDs, nil
}

// Get a match by match id.
func (c *Client) GetMatchByMatchID(matchId string) (*models.Match, error) {
	u := fmt.Sprintf("%s/lol/match/v5/matches/%s", c.routingServerUrl, matchId)

	var match models.Match
	resp, err := c.Get(u, &match)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get match by ID",
			Headers:    resp.Header,
		}
	}

	return &match, nil
}

// Get a match timeline by match id
func (c *Client) GetMatchTimelineByMatchID(matchId string) (*models.Timeline, error) {
	u := fmt.Sprintf("%s/lol/match/v5/matches/%s/timeline", c.routingServerUrl, matchId)

	var mtl models.Timeline
	resp, err := c.Get(u, &mtl)
	if err != nil {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to get match timeline by ID",
			Headers:    resp.Header,
		}
	}

	return &mtl, nil
}
