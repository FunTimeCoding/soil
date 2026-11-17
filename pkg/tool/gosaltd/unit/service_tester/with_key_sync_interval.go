package service_tester

import "time"

func WithKeySyncInterval(d time.Duration) Option {
	return func(o *Tester) {
		o.keySyncInterval = d
	}
}
