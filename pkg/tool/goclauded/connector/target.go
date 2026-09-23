package connector

import "time"

type Target struct {
	Identifier    string
	Name          string
	LastSeen      time.Time
	LastPromptAt  *time.Time
	LastTurnEndAt *time.Time
	ClosedAt      *time.Time
	Labels        map[string]string
}
