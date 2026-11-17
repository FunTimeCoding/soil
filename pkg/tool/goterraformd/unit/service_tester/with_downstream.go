package service_tester

import provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"

func WithDownstream(downstream ...*provisionMock.Client) Option {
	return func(o *Tester) {
		o.downstream = downstream
	}
}
