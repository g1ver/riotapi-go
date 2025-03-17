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

	fmt.Println(u.String())

	var matchIDs []string
	if err := c.Get(u.String(), &matchIDs); err != nil {
		return nil, fmt.Errorf("failed to get match IDs: %w", err)
	}

	return matchIDs, nil
}
