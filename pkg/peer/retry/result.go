package retry

import "time"

type Result struct {
	Attempts int
	Elapsed  time.Duration
}
