package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"

func WithNetbox() Option {
	return func(o *Tester) {
		o.Netbox = mock_client.New(o.t)
	}
}
