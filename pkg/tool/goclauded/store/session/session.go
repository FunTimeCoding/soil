package session

import "time"

type Session struct {
	Identifier          string    `gorm:"primaryKey;column:identifier"`
	Name                string    `gorm:"column:name"`
	Callsign            *string   `gorm:"uniqueIndex;column:callsign"`
	ModelContextSession string    `gorm:"column:model_context_session"`
	Topic               string    `gorm:"column:topic"`
	Files               string    `gorm:"column:files"`
	Alias               *string   `gorm:"uniqueIndex;column:alias"`
	Description         string    `gorm:"column:description"`
	Slug                string    `gorm:"column:slug"`
	FirstMessage        string    `gorm:"column:first_message"`
	TurnCount           int       `gorm:"column:turn_count"`
	Lines               int       `gorm:"column:lines"`
	TokenUsage          string    `gorm:"column:token_usage"`
	WorkDirectory       string    `gorm:"column:work_directory"`
	Branch              string    `gorm:"column:branch"`
	SessionTimestamp    string    `gorm:"column:session_timestamp"`
	Listening           bool      `gorm:"column:listening"`
	TimedOut            string    `gorm:"column:timed_out"`
	LastSeen            time.Time `gorm:"column:last_seen"`
	StartedAt           time.Time `gorm:"column:started_at"`
	LastActiveAt        time.Time `gorm:"column:last_active_at"`
}
