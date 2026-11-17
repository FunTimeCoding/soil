package outage_tracker

import "time"

func (r *Tracker) WithClock(clock func() time.Time) *Tracker {
	r.clock = clock

	return r
}
