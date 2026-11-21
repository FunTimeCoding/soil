package score_result

import "time"

type Result struct {
	Filled   bool
	Mirrored int
	Start    int
	Total    int
	Synced   time.Time
}
