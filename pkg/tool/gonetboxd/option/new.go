package option

import "github.com/funtimecoding/soil/pkg/netbox"

func New() *Netbox {
	return &Netbox{Client: netbox.NewEnvironment()}
}
