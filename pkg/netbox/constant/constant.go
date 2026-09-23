package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"regexp"
)

const (
	HostEnvironment  = "NET_BOX_HOST"
	TokenEnvironment = "NET_BOX_TOKEN" // #nosec G101 not a hardcoded secret

	NoName           = "no name"
	NoGroup          = "no group"
	NoTenant         = "no tenant"
	NoPrimaryAddress = "no primary address"
	NoComment        = "no comment"
	NoObjectType     = "no object type"
	NoDevice         = "no device"
	NoSerial         = "no serial"
	NoType           = "no type"

	PageLimit int32 = 1000

	DeviceAddress           = "dcim.device"
	InterfaceAddress        = "dcim.interface"
	VirtualInterfaceAddress = "virtualization.vminterface"
	VirtualMachineAddress   = "virtualization.virtualmachine"

	PrefixAddress = "ipam.prefix"

	Interface = "/api"

	SignatureHeader = "X-Hook-Signature"

	FixtureAddress         = "192.168.0.1/24"
	FixturePhysicalAddress = "02:00:00:00:00:01"

	DeviceActiveStatus = "active"
)

var (
	Format = constant.ColorFormat.Copy()

	InternetAddressTargets = []string{
		DeviceAddress,
		InterfaceAddress,
		VirtualMachineAddress,
		VirtualInterfaceAddress,
		PrefixAddress,
	}

	PhysicalAddressTargets = []string{InterfaceAddress, VirtualInterfaceAddress}
)

var NonSlug = regexp.MustCompile(`[^a-z0-9-]`)
