package device

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func (d *Device) LookupHost() string {
	if d.PrimaryAddress == constant.NoPrimaryAddress {
		return ""
	}

	return system.LookupFirst(d.PrimaryAddress)
}
