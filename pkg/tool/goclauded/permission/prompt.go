package permission

import "time"

type Prompt struct {
	Session  string
	When     time.Time
	Command  string
	Reason   string
	Declined bool
	Cleared  bool
	Blockers []string
}
