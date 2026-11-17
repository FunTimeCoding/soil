package service_tester

func WithReferenceRoot(root string) Option {
	return func(o *Tester) {
		o.referenceRoot = root
	}
}
