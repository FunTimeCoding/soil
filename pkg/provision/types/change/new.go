package change

import "github.com/funtimecoding/soil/pkg/provision/types/host_event"

func New(
	addresses []string,
	hosts []*host_event.Event,
) *Change {
	return &Change{Addresses: addresses, Hosts: hosts}
}
