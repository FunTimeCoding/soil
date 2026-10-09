package reacher

import "time"

func (r *Reacher) WithClock(clock func() time.Time) *Reacher {
	r.clock = clock

	return r
}
