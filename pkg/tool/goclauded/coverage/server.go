package coverage

import "time"

type Server struct {
	Name        string
	Path        string
	Configured  bool
	Registered  int
	UsedTotal   int
	UsedRecent  int
	CallsTotal  int
	CallsRecent int
	LastUsed    time.Time
	Tools       []*Tool
}
