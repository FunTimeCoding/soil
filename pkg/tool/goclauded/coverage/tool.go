package coverage

import "time"

type Tool struct {
	Name        string
	Registered  bool
	CallsTotal  int
	CallsRecent int
	LastUsed    time.Time
}
