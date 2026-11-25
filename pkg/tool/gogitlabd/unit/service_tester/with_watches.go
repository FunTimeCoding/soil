package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"

func WithWatches() Option {
	return func(o *Tester) {
		o.connector = mock_client.New(o.t)
	}
}
