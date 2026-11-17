package service_tester

func WithProject(project []string) Option {
	return func(o *Tester) {
		o.project = project
	}
}
