package service_tester

import "time"

func WithClock(at time.Time) Option {
	return func(s *Tester) {
		s.now = func() time.Time { return at }
	}
}
