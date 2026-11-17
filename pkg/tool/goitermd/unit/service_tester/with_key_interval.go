package service_tester

import "time"

func WithKeyInterval(d time.Duration) Option {
	return func(o *Tester) {
		o.keyInterval = d
	}
}
