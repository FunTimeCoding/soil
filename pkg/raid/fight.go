package raid

import "time"

type Fight struct {
	Filename             string    `gorm:"primaryKey" json:"filename"`
	Timestamp            time.Time `json:"timestamp"`
	DurationMS           int       `json:"duration_ms"`
	MapIdentifier        int       `gorm:"column:map_id" json:"map_id"`
	MapName              string    `json:"map_name"`
	AlliedCount          int       `json:"allied_count"`
	AlliedTeamIdentifier int       `gorm:"column:allied_team_id" json:"allied_team_id"`
	EnemyCount           int       `json:"enemy_count"`
	EnemyTeams           *string   `json:"enemy_teams,omitempty"`
	Success              bool      `json:"success"`
	Enriched             bool      `gorm:"not null;default:false" json:"enriched"`
	RaidIdentifier       *uint     `gorm:"column:raid_id" json:"raid_id,omitempty"`
}
