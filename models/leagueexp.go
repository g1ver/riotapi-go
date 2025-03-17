package models

type LeagueEntry struct {
	LeagueID     string     `json:"leagueId"`
	SummonerID   string     `json:"summonerId"`
	PUUID        string     `json:"puuid"`
	QueueType    string     `json:"queueType"`
	Tier         string     `json:"tier"`
	Rank         string     `json:"rank"`
	LeaguePoints int        `json:"leaguePoints"`
	Wins         int        `json:"wins"`
	Losses       int        `json:"losses"`
	HotStreak    bool       `json:"hotStreak"`
	Veteran      bool       `json:"veteran"`
	FreshBlood   bool       `json:"freshBlood"`
	Inactive     bool       `json:"inactive"`
	MiniSeries   MiniSeries `json:"miniSeries"`
}

type MiniSeries struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}
