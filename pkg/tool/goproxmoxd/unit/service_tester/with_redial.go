package service_tester

func WithRedial() Option {
	return func(o *Tester) {
		o.redial = true
	}
}
