package change

import "github.com/funtimecoding/soil/pkg/provision/types/host_event"

type Change struct {
	Addresses []string
	Hosts     []*host_event.Event
}
