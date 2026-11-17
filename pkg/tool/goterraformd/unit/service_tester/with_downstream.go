package service_tester

import "github.com/funtimecoding/soil/pkg/provision/client/mock_client"

func WithDownstream(downstream ...*mock_client.Client) Option {
	return func(o *Tester) {
		o.downstream = downstream
	}
}
