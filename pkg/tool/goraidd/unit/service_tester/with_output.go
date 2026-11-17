package service_tester

func WithOutput(path string) Option {
	return func(o *Tester) {
		o.Output = path
	}
}
