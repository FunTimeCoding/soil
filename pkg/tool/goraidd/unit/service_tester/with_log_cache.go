package service_tester

func WithLogCache(path string) Option {
	return func(o *Tester) {
		o.LogCache = path
	}
}
