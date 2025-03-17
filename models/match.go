package models

type GetMatchIDsByPUUIDParams struct {
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
	Queue     int    `url:"queue,omitempty"`
	Type      string `url:"type,omitempty"`
	Start     int    `url:"start,omitempty"`
	Count     int    `url:"count,omitempty"`
}

type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type Info struct {
	EndOfGameResult    string        `json:"endOfGameResult"`
	GameCreation       int64         `json:"gameCreation"`
	GameDuration       int64         `json:"gameDuration"`
	GameEndTimestamp   int64         `json:"gameEndTimestamp"`
	GameID             int64         `json:"gameId"`
	GameMode           string        `json:"gameMode"`
	GameName           string        `json:"gameName"`
	GameStartTimestamp int64         `json:"gameStartTimestamp"`
	GameType           string        `json:"gameType"`
	GameVersion        string        `json:"gameVersion"`
	MapID              int           `json:"mapId"`
	Participants       []Participant `json:"participants"`
	PlatformID         string        `json:"platformId"`
	QueueID            int           `json:"queueId"`
	Teams              []ClashTeam   `json:"teams"`
	TournamentCode     string        `json:"tournamentCode"`
}

type Participant struct {
	AllInPings                     int        `json:"allInPings"`
	AssistMePings                  int        `json:"assistMePings"`
	Assists                        int        `json:"assists"`
	BaronKills                     int        `json:"baronKills"`
	BountyLevel                    int        `json:"bountyLevel"`
	ChampExperience                int        `json:"champExperience"`
	ChampLevel                     int        `json:"champLevel"`
	ChampionID                     int        `json:"championId"`
	ChampionName                   string     `json:"championName"`
	CommandPings                   int        `json:"commandPings"`
	ChampionTransform              int        `json:"championTransform"`
	ConsumablesPurchased           int        `json:"consumablesPurchased"`
	Challenges                     Challenges `json:"challenges"`
	DamageDealtToBuildings         int        `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int        `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int        `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int        `json:"damageSelfMitigated"`
	Deaths                         int        `json:"deaths"`
	DetectorWardsPlaced            int        `json:"detectorWardsPlaced"`
	DoubleKills                    int        `json:"doubleKills"`
	DragonKills                    int        `json:"dragonKills"`
	EligibleForProgression         bool       `json:"eligibleForProgression"`
	EnemyMissingPings              int        `json:"enemyMissingPings"`
	EnemyVisionPings               int        `json:"enemyVisionPings"`
	FirstBloodAssist               bool       `json:"firstBloodAssist"`
	FirstBloodKill                 bool       `json:"firstBloodKill"`
	FirstTowerAssist               bool       `json:"firstTowerAssist"`
	FirstTowerKill                 bool       `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool       `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool       `json:"gameEndedInSurrender"`
	HoldPings                      int        `json:"holdPings"`
	GetBackPings                   int        `json:"getBackPings"`
	GoldEarned                     int        `json:"goldEarned"`
	GoldSpent                      int        `json:"goldSpent"`
	IndividualPosition             string     `json:"individualPosition"`
	InhibitorKills                 int        `json:"inhibitorKills"`
	InhibitorTakedowns             int        `json:"inhibitorTakedowns"`
	InhibitorsLost                 int        `json:"inhibitorsLost"`
	Item0                          int        `json:"item0"`
	Item1                          int        `json:"item1"`
	Item2                          int        `json:"item2"`
	Item3                          int        `json:"item3"`
	Item4                          int        `json:"item4"`
	Item5                          int        `json:"item5"`
	Item6                          int        `json:"item6"`
	ItemsPurchased                 int        `json:"itemsPurchased"`
	KillingSprees                  int        `json:"killingSprees"`
	Kills                          int        `json:"kills"`
	Lane                           string     `json:"lane"`
	LargestCriticalStrike          int        `json:"largestCriticalStrike"`
	LargestKillingSpree            int        `json:"largestKillingSpree"`
	LargestMultiKill               int        `json:"largestMultiKill"`
	LongestTimeSpentLiving         int        `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int        `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int        `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int        `json:"magicDamageTaken"`
	Missions                       Missions   `json:"missions"`
	NeutralMinionsKilled           int        `json:"neutralMinionsKilled"`
	NeedVisionPings                int        `json:"needVisionPings"`
	NexusKills                     int        `json:"nexusKills"`
	NexusTakedowns                 int        `json:"nexusTakedowns"`
	NexusLost                      int        `json:"nexusLost"`
	ObjectivesStolen               int        `json:"objectivesStolen"`
	ObjectivesStolenAssists        int        `json:"objectivesStolenAssists"`
	OnMyWayPings                   int        `json:"onMyWayPings"`
	ParticipantID                  int        `json:"participantId"`
	PlayerScore0                   int        `json:"playerScore0"`
	PlayerScore1                   int        `json:"playerScore1"`
	PlayerScore2                   int        `json:"playerScore2"`
	PlayerScore3                   int        `json:"playerScore3"`
	PlayerScore4                   int        `json:"playerScore4"`
	PlayerScore5                   int        `json:"playerScore5"`
	PlayerScore6                   int        `json:"playerScore6"`
	PlayerScore7                   int        `json:"playerScore7"`
	PlayerScore8                   int        `json:"playerScore8"`
	PlayerScore9                   int        `json:"playerScore9"`
	PlayerScore10                  int        `json:"playerScore10"`
	PlayerScore11                  int        `json:"playerScore11"`
	PentaKills                     int        `json:"pentaKills"`
	Perks                          Perks      `json:"perks"`
	PhysicalDamageDealt            int        `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int        `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int        `json:"physicalDamageTaken"`
	Placement                      int        `json:"placement"`
	PlayerAugment1                 int        `json:"playerAugment1"`
	PlayerAugment2                 int        `json:"playerAugment2"`
	PlayerAugment3                 int        `json:"playerAugment3"`
	PlayerAugment4                 int        `json:"playerAugment4"`
	PlayerSubteamID                int        `json:"playerSubteamId"`
	PushPings                      int        `json:"pushPings"`
	ProfileIcon                    int        `json:"profileIcon"`
	PUUID                          string     `json:"puuid"`
	QuadraKills                    int        `json:"quadraKills"`
	RiotIDGameName                 string     `json:"riotIdGameName"`
	RiotIDTagline                  string     `json:"riotIdTagline"`
	Role                           string     `json:"role"`
	SightWardsBoughtInGame         int        `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int        `json:"spell1Casts"`
	Spell2Casts                    int        `json:"spell2Casts"`
	Spell3Casts                    int        `json:"spell3Casts"`
	Spell4Casts                    int        `json:"spell4Casts"`
	SubteamPlacement               int        `json:"subteamPlacement"`
	Summoner1Casts                 int        `json:"summoner1Casts"`
	Summoner1ID                    int        `json:"summoner1Id"`
	Summoner2Casts                 int        `json:"summoner2Casts"`
	Summoner2ID                    int        `json:"summoner2Id"`
	SummonerID                     string     `json:"summonerId"`
	SummonerLevel                  int        `json:"summonerLevel"`
	SummonerName                   string     `json:"summonerName"`
	TeamEarlySurrendered           bool       `json:"teamEarlySurrendered"`
	TeamID                         int        `json:"teamId"`
	TeamPosition                   string     `json:"teamPosition"`
	TimeCCingOthers                int        `json:"timeCCingOthers"`
	TimePlayed                     int        `json:"timePlayed"`
	TotalAllyJungleMinionsKilled   int        `json:"totalAllyJungleMinionsKilled"`
	TotalDamageDealt               int        `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int        `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int        `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int        `json:"totalDamageTaken"`
	TotalEnemyJungleMinionsKilled  int        `json:"totalEnemyJungleMinionsKilled"`
	TotalHeal                      int        `json:"totalHeal"`
	TotalHealsOnTeammates          int        `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int        `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int        `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int        `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int        `json:"totalUnitsHealed"`
	TripleKills                    int        `json:"tripleKills"`
	TrueDamageDealt                int        `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int        `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int        `json:"trueDamageTaken"`
	TurretKills                    int        `json:"turretKills"`
	TurretTakedowns                int        `json:"turretTakedowns"`
	TurretsLost                    int        `json:"turretsLost"`
	UnrealKills                    int        `json:"unrealKills"`
	VisionScore                    int        `json:"visionScore"`
	VisionClearedPings             int        `json:"visionClearedPings"`
	VisionWardsBoughtInGame        int        `json:"visionWardsBoughtInGame"`
	WardsKilled                    int        `json:"wardsKilled"`
	WardsPlaced                    int        `json:"wardsPlaced"`
	Win                            bool       `json:"win"`
}

type Challenges struct {
	AssistStreakCount                         int     `json:"12AssistStreakCount"`
	BaronBuffGoldAdvantageOverThreshold       int     `json:"baronBuffGoldAdvantageOverThreshold"`
	ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
	EarliestBaron                             float64 `json:"earliestBaron"`
	EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
	EarliestElderDragon                       float64 `json:"earliestElderDragon"`
	EarlyLaningPhaseGoldExpAdvantage          float64 `json:"earlyLaningPhaseGoldExpAdvantage"`
	FasterSupportQuestCompletion              int     `json:"fasterSupportQuestCompletion"`
	FastestLegendary                          int     `json:"fastestLegendary"`
	HadAfkTeammate                            int     `json:"hadAfkTeammate"`
	HighestChampionDamage                     int     `json:"highestChampionDamage"`
	HighestCrowdControlScore                  int     `json:"highestCrowdControlScore"`
	HighestWardKills                          int     `json:"highestWardKills"`
	JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle"`
	KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler"`
	LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                            int     `json:"legendaryCount"`
	MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
	MostWardsDestroyedOneSweeper              int     `json:"mostWardsDestroyedOneSweeper"`
	MythicItemUsed                            int     `json:"mythicItemUsed"`
	PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
	SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
	TakedownsFirst25Minutes                   int     `json:"takedownsFirst25Minutes"`
	TeleportTakedowns                         int     `json:"teleportTakedowns"`
	ThirdInhibitorDestroyedTime               float64 `json:"thirdInhibitorDestroyedTime"`
	ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
	VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
	InfernalScalePickup                       int     `json:"InfernalScalePickup"`
	FistBumpParticipation                     int     `json:"fistBumpParticipation"`
	VoidMonsterKill                           int     `json:"voidMonsterKill"`
	AbilityUses                               int     `json:"abilityUses"`
	AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                  float64 `json:"alliedJungleMonsterKills"`
	BaronTakedowns                            int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                                float64 `json:"bountyGold"`
	BuffsStolen                               int     `json:"buffsStolen"`
	CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
	ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
	DamagePerMinute                           float64 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                                int     `json:"doubleAces"`
	DragonTakedowns                           int     `json:"dragonTakedowns"`
	LegendaryItemUsed                         []int   `json:"legendaryItemUsed"`
	EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                   float64 `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
	EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
	FirstTurretKilled                         int     `json:"firstTurretKilled"`
	FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
	FlawlessAces                              int     `json:"flawlessAces"`
	FullTeamTakedown                          int     `json:"fullTeamTakedown"`
	GameLength                                float64 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                             float64 `json:"goldPerMinute"`
	HadOpenNexus                              int     `json:"hadOpenNexus"`
	ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
	InitialBuffCount                          int     `json:"initialBuffCount"`
	InitialCrabCount                          int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
	JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	KDA                                       float64 `json:"kda"`
	KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
	KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
	KillingSprees                             int     `json:"killingSprees"`
	KillParticipation                         float64 `json:"killParticipation"`
	KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
	KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
	KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
	KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
	LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
	LostAnInhibitor                           int     `json:"lostAnInhibitor"`
	MaxKillDeficit                            int     `json:"maxKillDeficit"`
	MejaisFullStackInTime                     int     `json:"mejaisFullStackInTime"`
	MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
	Multikills                                int     `json:"multikills"`
	MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
	MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
	OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                          int     `json:"outnumberedKills"`
	OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                               int     `json:"perfectGame"`
	PickKillWithAlly                          int     `json:"pickKillWithAlly"`
	PoroExplosions                            int     `json:"poroExplosions"`
	QuickCleanse                              int     `json:"quickCleanse"`
	QuickFirstTurret                          int     `json:"quickFirstTurret"`
	QuickSoloKills                            int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
	SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
	ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
	ShortestTimeToAceFromFirstTakedown        float64 `json:"shortestTimeToAceFromFirstTakedown"`
	SkillshotsDodged                          int     `json:"skillshotsDodged"`
	SkillshotsHit                             int     `json:"skillshotsHit"`
	SnowballsHit                              int     `json:"snowballsHit"`
	SoloBaronKills                            int     `json:"soloBaronKills"`
	SWARM_DefeatAatrox                        int     `json:"SWARM_DefeatAatrox"`
	SWARM_DefeatBriar                         int     `json:"SWARM_DefeatBriar"`
	SWARM_DefeatMiniBosses                    int     `json:"SWARM_DefeatMiniBosses"`
	SWARM_EvolveWeapon                        int     `json:"SWARM_EvolveWeapon"`
	SWARM_Have3Passives                       int     `json:"SWARM_Have3Passives"`
	SWARM_KillEnemy                           int     `json:"SWARM_KillEnemy"`
	SWARM_PickupGold                          float64 `json:"SWARM_PickupGold"`
	SWARM_ReachLevel50                        int     `json:"SWARM_ReachLevel50"`
	SWARM_Survive15Min                        int     `json:"SWARM_Survive15Min"`
	SWARM_WinWith5EvolvedWeapons              int     `json:"SWARM_WinWith5EvolvedWeapons"`
	SoloKills                                 int     `json:"soloKills"`
	StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
	SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
	TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
	Takedowns                                 int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
	TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                            int     `json:"teamBaronKills"`
	TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
	TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
	TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
	TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
	TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
	TurretTakedowns                           int     `json:"turretTakedowns"`
	TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
	TwoWardsOneSweeperCount                   int     `json:"twoWardsOneSweeperCount"`
	UnseenRecalls                             int     `json:"unseenRecalls"`
	VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
	WardsGuarded                              int     `json:"wardsGuarded"`
	WardTakedowns                             int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
}

type Missions struct {
	PlayerScore0  int `json:"playerScore0"`
	PlayerScore1  int `json:"playerScore1"`
	PlayerScore2  int `json:"playerScore2"`
	PlayerScore3  int `json:"playerScore3"`
	PlayerScore4  int `json:"playerScore4"`
	PlayerScore5  int `json:"playerScore5"`
	PlayerScore6  int `json:"playerScore6"`
	PlayerScore7  int `json:"playerScore7"`
	PlayerScore8  int `json:"playerScore8"`
	PlayerScore9  int `json:"playerScore9"`
	PlayerScore10 int `json:"playerScore10"`
	PlayerScore11 int `json:"playerScore11"`
}

type Perks struct {
	StatPerks PerkStats   `json:"statPerks"`
	Styles    []PerkStyle `json:"styles"`
}

type PerkStats struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type PerkStyle struct {
	Description string               `json:"description"`
	Selections  []PerkStyleSelection `json:"selections"`
	Style       int                  `json:"style"`
}

type PerkStyleSelection struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

type Team struct {
	Bans       []Ban      `json:"bans"`
	Objectives Objectives `json:"objectives"`
	TeamID     int        `json:"teamId"`
	Win        bool       `json:"win"`
}

type Ban struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type Objectives struct {
	Baron      Objective `json:"baron"`
	Champion   Objective `json:"champion"`
	Dragon     Objective `json:"dragon"`
	Horde      Objective `json:"horde"`
	Inhibitor  Objective `json:"inhibitor"`
	RiftHerald Objective `json:"riftHerald"`
	Tower      Objective `json:"tower"`
}

type Objective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type MetadataTimeLine struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type InfoTimeLine struct {
	EndOfGameResult string                `json:"endOfGameResult"`
	FrameInterval   int64                 `json:"frameInterval"`
	GameID          int64                 `json:"gameId"`
	Participants    []ParticipantTimeLine `json:"participants"`
	Frames          []FramesTimeLine      `json:"frames"`
}

type ParticipantTimeLine struct {
	ParticipantID int    `json:"participantId"`
	PUUID         string `json:"puuid"`
}

type FramesTimeLine struct {
	Events            []EventsTimeLine            `json:"events"`
	ParticipantFrames map[string]ParticipantFrame `json:"participantFrames"`
	Timestamp         int                         `json:"timestamp"`
}

type EventsTimeLine struct {
	Timestamp     int64  `json:"timestamp"`
	RealTimestamp int64  `json:"realTimestamp"`
	Type          string `json:"type"`
}

type ParticipantFrame struct {
	ChampionStats            ChampionStats `json:"championStats"`
	CurrentGold              int           `json:"currentGold"`
	DamageStats              DamageStats   `json:"damageStats"`
	GoldPerSecond            int           `json:"goldPerSecond"`
	JungleMinionsKilled      int           `json:"jungleMinionsKilled"`
	Level                    int           `json:"level"`
	MinionsKilled            int           `json:"minionsKilled"`
	ParticipantID            int           `json:"participantId"`
	Position                 Position      `json:"position"`
	TimeEnemySpentControlled int           `json:"timeEnemySpentControlled"`
	TotalGold                int           `json:"totalGold"`
	XP                       int           `json:"xp"`
}

type ChampionStats struct {
	AbilityHaste         int `json:"abilityHaste"`
	AbilityPower         int `json:"abilityPower"`
	Armor                int `json:"armor"`
	ArmorPen             int `json:"armorPen"`
	ArmorPenPercent      int `json:"armorPenPercent"`
	AttackDamage         int `json:"attackDamage"`
	AttackSpeed          int `json:"attackSpeed"`
	BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
	BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
	CCReduction          int `json:"ccReduction"`
	CooldownReduction    int `json:"cooldownReduction"`
	Health               int `json:"health"`
	HealthMax            int `json:"healthMax"`
	HealthRegen          int `json:"healthRegen"`
	Lifesteal            int `json:"lifesteal"`
	MagicPen             int `json:"magicPen"`
	MagicPenPercent      int `json:"magicPenPercent"`
	MagicResist          int `json:"magicResist"`
	MovementSpeed        int `json:"movementSpeed"`
	Omnivamp             int `json:"omnivamp"`
	PhysicalVamp         int `json:"physicalVamp"`
	Power                int `json:"power"`
	PowerMax             int `json:"powerMax"`
	PowerRegen           int `json:"powerRegen"`
	SpellVamp            int `json:"spellVamp"`
}

type DamageStats struct {
	MagicDamageDone               int `json:"magicDamageDone"`
	MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
	MagicDamageTaken              int `json:"magicDamageTaken"`
	PhysicalDamageDone            int `json:"physicalDamageDone"`
	PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
	PhysicalDamageTaken           int `json:"physicalDamageTaken"`
	TotalDamageDone               int `json:"totalDamageDone"`
	TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
	TotalDamageTaken              int `json:"totalDamageTaken"`
	TrueDamageDone                int `json:"trueDamageDone"`
	TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
	TrueDamageTaken               int `json:"trueDamageTaken"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Timeline struct {
	Metadata MetadataTimeLine `json:"metadata"`
	Info     InfoTimeLine     `json:"info"`
}
