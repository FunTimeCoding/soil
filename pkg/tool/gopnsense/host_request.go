package gopnsense

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"

func hostRequest(f *hostFlags) *client.HostRequest {
	return &client.HostRequest{
		Host:             optional(f.host),
		Domain:           optional(f.domain),
		Address:          optional(f.address),
		HardwareAddress:  optional(f.hardwareAddress),
		ClientIdentifier: optional(f.clientIdentifier),
		Description:      optional(f.description),
	}
}
