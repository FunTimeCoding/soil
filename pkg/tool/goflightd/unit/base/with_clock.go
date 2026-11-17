package base

import "time"

func WithClock(c func() time.Time) Option {
	return func(o *Server) {
		o.clock = c
	}
}
