package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

const (
	HostEnvironment  = "NET_BOX_HOST"
	TokenEnvironment = "NET_BOX_TOKEN" // #nosec G101 not a hardcoded secret

	NoName           = "no name"
	NoGroup          = "no group"
	NoTenant         = "no tenant"
	NoPrimaryAddress = "no primary address"
	NoComment        = "no comment"

	PageLimit int32 = 1000

	DeviceAddress           = "dcim.device"
	InterfaceAddress        = "dcim.interface"
	VirtualInterfaceAddress = "virtualization.vminterface"
	VirtualMachineAddress   = "virtualization.virtualmachine"

	Interface = "/api"

	SignatureHeader = "X-Hook-Signature"

	FixtureAddress = "192.168.0.1/24"
)

var (
	Format        = option.Color.Copy()
	ErrorNotFound = errors.New("not found")
)
