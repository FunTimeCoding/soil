package cruise

import "time"

type Reading struct {
	Closed     bool
	StateKnown bool
	Idle       bool
	IdleSince  time.Time
	Grace      time.Duration
}
