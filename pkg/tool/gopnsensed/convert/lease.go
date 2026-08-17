package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/lease"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Lease(v *lease.Lease) *server.Lease {
	return &server.Lease{
		Address:          v.Address,
		HardwareAddress:  v.HardwareAddress,
		Hostname:         v.Hostname,
		ClientIdentifier: v.ClientIdentifier,
		Expire:           v.Expire,
		Interface:        v.Interface,
		Reserved:         v.Reserved,
	}
}
