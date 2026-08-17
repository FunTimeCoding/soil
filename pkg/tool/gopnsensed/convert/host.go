package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/host"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Host(v *host.Host) *server.Host {
	return &server.Host{
		Identifier:       v.Identifier,
		Host:             v.Host,
		Domain:           v.Domain,
		Address:          v.Address,
		HardwareAddress:  v.HardwareAddress,
		ClientIdentifier: v.ClientIdentifier,
		Description:      v.Description,
	}
}
