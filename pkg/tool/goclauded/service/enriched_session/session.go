package enriched_session

import "time"

type Session struct {
	Identifier    string
	Slug          string
	Timestamp     string
	LastSeen      time.Time
	WorkDirectory string
	Branch        string
	Lines         int
	Name          string
	Alias         string
	Description   string
	TurnCount     int
	Active        bool
}
