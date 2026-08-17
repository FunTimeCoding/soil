package lease

import (
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"time"
)

func New(v response.Lease) *Lease {
	return &Lease{
		Address:          v.Address,
		HardwareAddress:  v.HardwareAddress,
		Hostname:         v.Hostname,
		ClientIdentifier: v.ClientIdentifier,
		Expire:           time.Unix(v.Expire, 0),
		Interface:        v.InterfaceName,
		Reserved:         len(v.Reserved) > 0,
	}
}
