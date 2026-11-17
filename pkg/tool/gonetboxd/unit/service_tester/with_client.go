package service_tester

import "github.com/funtimecoding/soil/pkg/netbox/face"

func WithClient(c face.NetBoxClient) Option {
	return func(o *Tester) {
		o.client = c
	}
}
