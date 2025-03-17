package models

type Content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type Update struct {
	ID               int       `json:"id"`
	Author           string    `json:"author"`
	Publish          bool      `json:"publish"`
	PublishLocations []string  `json:"publish_locations"`
	Translations     []Content `json:"translations"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
}

type Status struct {
	ID                int       `json:"id"`
	MaintenanceStatus string    `json:"maintenance_status"`
	IncidentSeverity  string    `json:"incident_severity"`
	Titles            []Content `json:"titles"`
	Updates           []Update  `json:"updates"`
	CreatedAt         string    `json:"created_at"`
	ArchiveAt         string    `json:"archive_at"`
	UpdatedAt         string    `json:"updated_at"`
	Platforms         []string  `json:"platforms"`
}

type PlatformData struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Locales      []string `json:"locales"`
	Maintenances []Status `json:"maintenances"`
	Incidents    []Status `json:"incidents"`
}
