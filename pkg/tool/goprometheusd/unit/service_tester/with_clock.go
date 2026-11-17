package service_tester

import "time"

func WithClock(c func() time.Time) Option {
	return func(o *Tester) {
		o.Clock = c
	}
}
