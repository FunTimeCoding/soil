package option

import "github.com/funtimecoding/soil/pkg/netbox"

type Netbox struct {
	Client          *netbox.Client
	Address         string
	ServiceTokens   []string
	LitePath        string
	PostgresLocator string
}
