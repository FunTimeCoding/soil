package host

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.Host) *Host {
	return &Host{
		Identifier:       v.Identifier,
		Host:             v.Host,
		Domain:           v.Domain,
		Address:          v.Address,
		HardwareAddress:  v.HardwareAddress,
		ClientIdentifier: v.ClientIdentifier,
		Description:      v.Description,
	}
}
