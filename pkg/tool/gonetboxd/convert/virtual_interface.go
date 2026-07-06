package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/netbox-community/go-netbox/v4"
)

func VirtualInterface(i *netbox.VMInterface) *server.VirtualInterface {
	return &server.VirtualInterface{
		Identifier: i.GetId(),
		Name:       i.GetName(),
	}
}
