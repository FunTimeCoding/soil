package service_tester

func WithPath(path string) Option {
	return func(o *Tester) {
		o.Path = path
	}
}
