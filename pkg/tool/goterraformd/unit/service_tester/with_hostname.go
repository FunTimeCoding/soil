package service_tester

func WithHostname(hostname string) Option {
	return func(o *Tester) {
		o.hostname = hostname
	}
}
