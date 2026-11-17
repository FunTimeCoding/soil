package base

import "time"

func WithHold(d time.Duration) Option {
	return func(s *setup) {
		s.hold = d
	}
}
