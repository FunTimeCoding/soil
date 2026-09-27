package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/netbox-community/go-netbox/v4"
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

	ObjectLinkField  = "url"
	InterfaceSegment = "/api/"

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

	Eth0 = "eth0"
	Eth1 = "eth1"

	InterfaceVirtual      = "virtual"
	InterfaceFastEthernet = "100base-tx"
	Interface1000BaseT    = "1000base-t"
	Interface2500BaseT    = "2.5gbase-t"

	// Status label
	RackActiveLabel     = "Active"
	RackDeprecatedLabel = "Deprecated"
	// Status value
	RackActive     = "active"
	RackDeprecated = "deprecated"
	RackUnknown    = "unknown"
	RackUnexpected = "unexpected"
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

	NonSlug = regexp.MustCompile(`[^a-z0-9-]`)

	InterfaceTypes = []netbox.InterfaceTypeValue{
		InterfaceVirtual,
		InterfaceFastEthernet,
		Interface1000BaseT,
		Interface2500BaseT,
	}

	ObjectTypeAlias = map[string]string{
		DeviceAddress:            "Device",
		InterfaceAddress:         "Interface",
		PrefixAddress:            "Prefix",
		VirtualInterfaceAddress:  "Interface",
		VirtualMachineAddress:    "Virtual Machine",
		"dcim.consoleport":       "Console Port",
		"dcim.devicerole":        "Device Role",
		"dcim.devicetype":        "Device Type",
		"dcim.location":          "Location",
		"dcim.moduletype":        "Module Type",
		"dcim.powerport":         "Power Port",
		"dcim.rack":              "Rack",
		"dcim.site":              "Site",
		"ipam.ipaddress":         "IP Address",
		"ipam.iprange":           "IP Range",
		"ipam.vlan":              "VLAN",
		"tenancy.tenant":         "Tenant",
		"virtualization.cluster": "Cluster",
	}
)
