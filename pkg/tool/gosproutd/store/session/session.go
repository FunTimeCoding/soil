package session

import "time"

type Session struct {
	Name    string
	Open    int
	Pending int
	Cleared int
	Bumped  int
	LastAt  time.Time
}
