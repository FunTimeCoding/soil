package unavailable

import "time"

type UnavailableError struct {
	Message string
	Since   time.Time
	Reason  string
}
