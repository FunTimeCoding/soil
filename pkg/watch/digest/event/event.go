package event

import "time"

type Event struct {
	Noun   string
	Author string
	Line   string
	At     time.Time
}
