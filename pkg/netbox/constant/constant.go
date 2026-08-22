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

	// SubnetAddress when assigning an IP address to a prefix/subnet
	SubnetAddress = "ipam.prefix"

	Interface = "/api"

	SignatureHeader = "X-Hook-Signature"

	FixtureAddress = "192.168.0.1/24"

	DeviceActiveStatus = "active"
)

var (
	Format = constant.ColorFormat.Copy()

	// InternetObjectTypes an IP address can be assigned to
	InternetObjectTypes = []string{
		DeviceAddress,
		InterfaceAddress,
		VirtualMachineAddress,
		VirtualInterfaceAddress,
		SubnetAddress,
	}

	// PhysicalObjectTypes a MAC address can be assigned to
	PhysicalObjectTypes = []string{InterfaceAddress}
)

var NonSlug = regexp.MustCompile(`[^a-z0-9-]`)
