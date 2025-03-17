package models

type ChampionMastery struct {
	PUUID                        string              `json:"puuid"`
	ChampionPointsUntilNextLevel int64               `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool                `json:"chestGranted"`
	ChampionID                   int64               `json:"championId"`
	LastPlayTime                 int64               `json:"lastPlayTime"`
	ChampionLevel                int                 `json:"championLevel"`
	ChampionPoints               int                 `json:"championPoints"`
	ChampionPointsSinceLastLevel int64               `json:"championPointsSinceLastLevel"`
	MarkRequiredForNextLevel     int                 `json:"markRequiredForNextLevel"`
	ChampionSeasonMilestone      int                 `json:"championSeasonMilestone"`
	NextSeasonMilestone          NextSeasonMilestone `json:"nextSeasonMilestone"`
	TokensEarned                 int                 `json:"tokensEarned"`
	MilestoneGrades              []string            `json:"milestoneGrades"`
}

type NextSeasonMilestone struct {
	RequireGradeCounts map[string]int `json:"requireGradeCounts"`
	RewardMarks        int            `json:"rewardMarks"`
	Bonus              bool           `json:"bonus"`
	RewardConfig       RewardConfig   `json:"rewardConfig"`
}

type RewardConfig struct {
	RewardValue   string `json:"rewardValue"`
	RewardType    string `json:"rewardType"`
	MaximumReward int    `json:"maximumReward"`
}
