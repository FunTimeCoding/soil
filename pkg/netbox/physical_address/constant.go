package physical_address

import "github.com/funtimecoding/soil/pkg/netbox/constant"

const (
	NoDevice     = "no device"
	NoObjectType = "no object type"
)

var ObjectTypes = []string{
	constant.InterfaceAddress,
}
