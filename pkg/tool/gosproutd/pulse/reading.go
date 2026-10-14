package pulse

import "time"

type Reading struct {
	Unpulsed int
	OldestAt time.Time
	NewestAt time.Time
	Closed   bool
}
