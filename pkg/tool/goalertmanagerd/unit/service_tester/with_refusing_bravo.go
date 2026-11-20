package service_tester

import "github.com/funtimecoding/soil/pkg/errors/validation"

func WithRefusingBravo() Option {
	return func(o *Tester) {
		WithInstances("alfa", "bravo")(o)
		o.Bravo().Alertmanager.Fail(validation.New("bad request"))
	}
}
