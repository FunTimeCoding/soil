package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func HostRequest(v *server.HostRequest) *request.Host {
	result := request.New()
	result.Host = text(v.Host)
	result.Domain = text(v.Domain)
	result.Address = text(v.Address)
	result.HardwareAddress = text(v.HardwareAddress)
	result.ClientIdentifier = text(v.ClientIdentifier)
	result.Description = text(v.Description)

	return result
}
