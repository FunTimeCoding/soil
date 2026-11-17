package service_tester

func WithPowerIntent() Option {
	return func(o *Tester) {
		o.powerIntent = true
	}
}
