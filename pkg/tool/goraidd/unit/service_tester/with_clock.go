package service_tester

import "time"

func WithClock(clock func() time.Time) Option {
	return func(o *Tester) {
		o.clock = clock
	}
}
