package models

type Summoner struct {
	AccountID     string `json:"accountId"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            string `json:"id"`
	PUUID         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
}
