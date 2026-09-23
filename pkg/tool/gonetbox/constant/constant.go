package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gonetbox",
	"NetBox infrastructure inventory CLI",
	"gonetbox [command]",
)

const VirtualPhysicalAddressUsage = "set-virtual-interface-physical-address [machine] [interface] [address]"
