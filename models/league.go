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

type LeagueItem struct {
	FreshBlood   bool       `json:"freshBlood"`
	Wins         int        `json:"wins"`
	MiniSeries   MiniSeries `json:"miniSeries"`
	Inactive     bool       `json:"inactive"`
	Veteran      bool       `json:"veteran"`
	HotStreak    bool       `json:"hotStreak"`
	Rank         string     `json:"rank"`
	LeaguePoints int        `json:"leaguePoints"`
	Losses       int        `json:"losses"`
	SummonerID   string     `json:"summonerId"`
	PUUID        string     `json:"puuid"`
}

type LeagueList struct {
	LeagueID string       `json:"leagueId"`
	Entries  []LeagueItem `json:"entries"`
	Tier     string       `json:"tier"`
	Name     string       `json:"name"`
	Queue    string       `json:"queue"`
}
