package option

import "github.com/funtimecoding/soil/pkg/netbox"

type Netbox struct {
	Client          *netbox.Client
	Address         string
	Version         string
	LitePath        string
	PostgresLocator string
}
