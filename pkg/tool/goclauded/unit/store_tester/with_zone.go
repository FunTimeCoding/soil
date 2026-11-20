package store_tester

import "time"

func WithZone(zone *time.Location) Option {
	return func(s *setup) {
		s.zone = zone
	}
}
