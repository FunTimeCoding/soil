package impression_call

import "time"

type Call struct {
	Content    string
	Source     string
	CapturedAt time.Time
}
