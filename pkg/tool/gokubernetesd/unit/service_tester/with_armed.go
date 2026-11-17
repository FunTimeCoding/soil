package service_tester

func WithArmed() Option {
	return func(o *Tester) {
		o.arm = true
	}
}
