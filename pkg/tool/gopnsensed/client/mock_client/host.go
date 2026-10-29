package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"

func host(
	identifier string,
	h client.HostRequest,
) client.Host {
	return client.Host{
		Identifier:       identifier,
		Host:             value(h.Host),
		Domain:           value(h.Domain),
		Address:          value(h.Address),
		HardwareAddress:  value(h.HardwareAddress),
		ClientIdentifier: value(h.ClientIdentifier),
		Description:      value(h.Description),
	}
}
