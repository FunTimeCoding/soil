package constant

import "time"

const (
	ContextKey  = "reachability"
	HostKey     = "host"
	StateKey    = "state"
	ReasonKey   = "reason"
	DurationKey = "duration"
	SinceKey    = "since"
	DownState   = "down"
	UpState     = "up"
)

const (
	Wait         = time.Minute
	Poll         = time.Second
	ProbeMinimum = 5 * time.Second
	ProbeMaximum = 5 * time.Minute
	Lasting      = 5 * time.Minute
)
