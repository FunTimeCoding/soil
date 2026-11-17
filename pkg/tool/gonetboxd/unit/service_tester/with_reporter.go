package service_tester

import "github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"

func WithReporter(r *memory.Memory) Option {
	return func(o *Tester) {
		o.Reporter = r
	}
}
