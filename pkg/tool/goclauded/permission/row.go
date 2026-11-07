package permission

import "time"

type Row struct {
	Blocker  string
	Prompts  int
	Sessions int
	Reason   string
	Last     time.Time
	Commands []string
}
