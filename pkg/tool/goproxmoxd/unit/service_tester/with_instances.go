package service_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
)

func WithInstances(names ...string) Option {
	return func(o *Tester) {
		o.instances = map[string]*Clients{}

		for _, n := range names {
			o.instances[n] = &Clients{
				Hypervisor: mock_client.New(o.t),
				Shell:      sshMock.New(o.t),
			}
		}
	}
}
