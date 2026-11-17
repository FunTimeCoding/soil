package service_tester

import "github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"

func WithUsage() Option {
	return func(o *Tester) {
		o.Usage = mock_client.New(o.t)
	}
}
