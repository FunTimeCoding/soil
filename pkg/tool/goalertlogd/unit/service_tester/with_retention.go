package service_tester

import "time"

func WithRetention(d time.Duration) Option {
	return func(o *Tester) {
		o.retention = d
	}
}
