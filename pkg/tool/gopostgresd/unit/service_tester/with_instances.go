package service_tester

import "github.com/funtimecoding/soil/pkg/postgres/mock_client"

func WithInstances(names ...string) Option {
	return func(o *Tester) {
		o.instances = map[string]*mock_client.Client{}

		for _, n := range names {
			o.instances[n] = mock_client.New(o.t)
		}
	}
}
