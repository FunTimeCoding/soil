package service_tester

func WithMonitor() Option {
	return func(o *Tester) {
		o.monitor = true
	}
}
