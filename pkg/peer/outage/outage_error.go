package outage

import "time"

type OutageError struct {
	Peer     string
	Since    time.Time
	Duration time.Duration
	cause    error
}
