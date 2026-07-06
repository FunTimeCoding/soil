package option

import "github.com/funtimecoding/soil/pkg/netbox"

type Netbox struct {
	Client  *netbox.Client
	Port    int
	Version string
}
