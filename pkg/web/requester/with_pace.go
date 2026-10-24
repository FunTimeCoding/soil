package requester

import (
	"golang.org/x/time/rate"
	"time"
)

func (r *Requester) WithPace(d time.Duration) *Requester {
	r.pace = rate.NewLimiter(rate.Every(d), 1)

	return r
}
