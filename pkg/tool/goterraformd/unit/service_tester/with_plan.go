package service_tester

func WithPlan(plan string) Option {
	return func(o *Tester) {
		o.plan = plan
	}
}
