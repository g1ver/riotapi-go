package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var regionBaseURLs = map[string]string{
	"br1":  "https://br1.api.riotgames.com",
	"eun1": "https://eun1.api.riotgames.com",
	"euw1": "https://euw1.api.riotgames.com",
	"jp1":  "https://jp1.api.riotgames.com",
	"kr":   "https://kr.api.riotgames.com",
	"la1":  "https://la1.api.riotgames.com",
	"la2":  "https://la2.api.riotgames.com",
	"me1":  "https://me1.api.riotgames.com",
	"na1":  "https://na1.api.riotgames.com",
	"oc1":  "https://oc1.api.riotgames.com",
	"ru":   "https://ru.api.riotgames.com",
	"sg2":  "https://sg2.api.riotgames.com",
	"tr1":  "https://tr1.api.riotgames.com",
	"tw2":  "https://tw2.api.riotgames.com",
	"vn2":  "https://vn2.api.riotgames.com",
}

// The AMERICAS routing value serves NA, BR, LAN and LAS. The ASIA routing value serves KR and JP. The EUROPE routing value serves EUNE, EUW, ME1, TR and RU. The SEA routing value serves OCE, SG2, TW2 and VN2.
var routingServers = map[string]string{
	"na1":  "americas",
	"br1":  "americas",
	"la1":  "americas",
	"la2":  "americas",
	"eun1": "europe",
	"euw1": "europe",
	"me1":  "europe",
	"tr1":  "europe",
	"ru":   "europe",
	"kr":   "asia",
	"jp1":  "asia",
	"oc1":  "sea",
	"sg2":  "sea",
	"tw2":  "sea",
	"vn2":  "sea",
}

type Client struct {
	httpClient    *http.Client
	apiKey        string
	region        string
	routingServer string
	baseUrl       string
}

func NewClient(apiKey string, region string) (*Client, error) {
	if _, exists := regionBaseURLs[region]; !exists {
		return nil, errors.New("Invalid region.")
	}

	if _, exists := routingServers[region]; !exists {
		return nil, errors.New("Invalid region.")
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey:        apiKey,
		region:        region,
		routingServer: routingServers[region],
		baseUrl:       regionBaseURLs[region],
	}, nil
}

func (c *Client) Get(url string, result any) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("X-Riot-Token", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
