package service_tester

func WithHiddenTag(tag string) Option {
	return func(o *Tester) {
		o.hiddenTag = tag
	}
}
