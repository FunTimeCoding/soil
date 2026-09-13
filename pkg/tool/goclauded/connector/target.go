package connector

import "time"

type Target struct {
	Identifier string
	Name       string
	LastSeen   time.Time
	Labels     map[string]string
}
