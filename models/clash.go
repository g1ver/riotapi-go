package models

type Player struct {
	SummonerId string `json:"summonerId"`
	PUUID      string `json:"puuid"`
	TeamId     string `json:"teamId"`
	Position   string `json:"position"`
	Role       string `json:"role"`
}

type Team struct {
	ID           string   `json:"id"`
	TournamentID int      `json:"tournamentId"`
	Name         string   `json:"name"`
	IconID       int      `json:"iconId"`
	Tier         int      `json:"tier"`
	Captain      string   `json:"captain"`
	Abbreviation string   `json:"abbreviation"`
	Players      []Player `json:"players"`
}

type Tournament struct {
	ID               int               `json:"id"`
	ThemeID          int               `json:"themeId"`
	NameKey          string            `json:"nameKey"`
	NameKeySecondary string            `json:"nameKeySecondary"`
	Schedule         []TournamentPhase `json:"schedule"`
}

type TournamentPhase struct {
	ID               int64 `json:"id"`
	RegistrationTime int64 `json:"registrationTime"`
	StartTime        int64 `json:"startTime"`
	Cancelled        bool  `json:"cancelled"`
}
