package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/face"

func WithConnector(k face.Client) Option {
	return func(o *Tester) {
		o.connector = k
	}
}
