package outage

import "time"

func New(
	peer string,
	since time.Time,
	duration time.Duration,
	cause error,
) *OutageError {
	return &OutageError{
		Peer:     peer,
		Since:    since,
		Duration: duration,
		cause:    cause,
	}
}
