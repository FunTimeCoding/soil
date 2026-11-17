package service_tester

import "github.com/funtimecoding/soil/pkg/provision/salt/face"

func WithMaster(master func() face.SaltClient) Option {
	return func(o *Tester) {
		o.master = master
	}
}
