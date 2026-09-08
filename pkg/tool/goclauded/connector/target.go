package connector

import "time"

type Target struct {
	Identifier string
	Name       string
	Timestamp  time.Time
	Labels     map[string]string
}
