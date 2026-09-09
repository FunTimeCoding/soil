package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type NetboxSource interface {
	AddDeviceJournalEntry(
		device string,
		kind string,
		comments string,
	) (*journal_entry.Entry, error)
	AddTag(
		deviceName string,
		tag string,
	) (*device.Device, error)
	AddVirtualJournalEntry(
		machine string,
		kind string,
		comments string,
	) (*journal_entry.Entry, error)
	AddVirtualTag(
		name string,
		tag string,
	) (*virtual_machine.Machine, error)
	Cables() ([]*cable.Cable, error)
	ClusterByName(n string) (*cluster.Cluster, error)
	Clusters() ([]*cluster.Cluster, error)
	ClusterTypeByName(n string) (*cluster_type.Type, error)
	ClusterTypes() ([]*cluster_type.Type, error)
	CreateAddress(
		interfaceIdentifier int32,
		address string,
		status string,
	) (*internet_address.Address, error)
	CreateCable(
		a *network.Interface,
		b *network.Interface,
	) (*cable.Cable, error)
	CreateCluster(
		name string,
		t *cluster_type.Type,
		s *site.Site,
	) (*cluster.Cluster, error)
	CreateClusterType(name string) (*cluster_type.Type, error)
	CreateDevice(
		name string,
		o *device_role.Role,
		tags []string,
		y *device_type.Type,
		s *site.Site,
		n *tenant.Tenant,
	) (*device.Device, error)
	CreateDeviceRole(name string) (*device_role.Role, error)
	CreateDeviceType(
		model string,
		m *manufacturer.Manufacturer,
	) (*device_type.Type, error)
	CreateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
	) (*network.Interface, error)
	CreateInternetAddressRange(
		start string,
		end string,
		status string,
		description string,
	) (*internet_address_range.Range, error)
	CreateLocation(
		name string,
		siteName string,
	) (*location.Location, error)
	CreateManufacturer(name string) (*manufacturer.Manufacturer, error)
	CreatePhysicalInterface(
		a net.HardwareAddr,
		description string,
		i *network.Interface,
	) (*physical_address.Address, error)
	CreatePlatform(name string) (*platform.Platform, error)
	CreatePrefix(
		cidr string,
		s *site.Site,
		description string,
	) (*prefix.Prefix, error)
	CreateSite(name string) (*site.Site, error)
	CreateTag(name string) (*tag.Tag, error)
	CreateTenant(name string) (*tenant.Tenant, error)
	CreateTunnel(
		name string,
		encapsulation string,
		group *tunnel_group.Group,
	) (*tunnel.Tunnel, error)
	CreateTunnelGroup(name string) (*tunnel_group.Group, error)
	CreateTunnelTermination(
		t *tunnel.Tunnel,
		terminationType string,
		terminationIdentifier int64,
		role string,
	) (*tunnel_termination.Termination, error)
	CreateVirtualAddress(
		interfaceIdentifier int32,
		address string,
		status string,
	) (*internet_address.Address, error)
	CreateVirtualDisk(
		machine string,
		name string,
		size int32,
	) (*virtual_disk.Disk, error)
	CreateVirtualInterface(
		vm *virtual_machine.Machine,
		name string,
	) (*netbox.VMInterface, error)
	CreateVirtualMachine(
		name string,
		cl *cluster.Cluster,
	) (*virtual_machine.Machine, error)
	CreateWirelessNetwork(ssid string) (*wireless_network.Network, error)
	DeleteInternet(identifier int32) error
	DeleteJournalEntry(identifier int32) error
	DeviceAddresses(device string) ([]*internet_address.Address, error)
	DeviceByName(n string) (*device.Device, error)
	DeviceInterfaceByName(
		d *device.Device,
		name string,
	) (*network.Interface, error)
	DeviceInterfaces(device int32) ([]*network.Interface, error)
	DeviceJournalEntries(
		device string,
		limit int32,
		offset int32,
	) ([]*journal_entry.Entry, error)
	DeviceRoleByName(n string) (*device_role.Role, error)
	DeviceRoles() ([]*device_role.Role, error)
	Devices() ([]*device.Device, error)
	DevicesByMatch(s string) ([]*device.Device, error)
	DeviceTagNames(name string) ([]string, error)
	DeviceTypeByName(n string) (*device_type.Type, error)
	DeviceTypes() ([]*device_type.Type, error)
	InternetAddressRanges() ([]*internet_address_range.Range, error)
	Locations() ([]*location.Location, error)
	ManufacturerByName(n string) (*manufacturer.Manufacturer, error)
	Manufacturers() ([]*manufacturer.Manufacturer, error)
	Platforms() ([]*platform.Platform, error)
	Prefixes() ([]*prefix.Prefix, error)
	RemoveTag(
		deviceName string,
		tag string,
	) (*device.Device, error)
	RemoveVirtualTag(
		name string,
		tag string,
	) (*virtual_machine.Machine, error)
	RenameDevice(
		name string,
		newName string,
	) (*device.Device, error)
	RenameVirtualMachine(
		name string,
		newName string,
	) (*virtual_machine.Machine, error)
	SetDeviceDescription(
		name string,
		description string,
	) (*device.Device, error)
	SetDeviceLocation(
		name string,
		locationName string,
	) (*device.Device, error)
	SetDevicePlatform(
		name string,
		platformName string,
	) (*device.Device, error)
	SetDevicePrimaryAddress(
		name string,
		address string,
	) (*device.Device, error)
	SetDeviceSerial(
		name string,
		serial string,
	) (*device.Device, error)
	SetDeviceStatus(
		name string,
		status string,
	) (*device.Device, error)
	SetDeviceTenant(
		name string,
		tenantName string,
	) (*device.Device, error)
	SetVirtualMachineCores(
		name string,
		cores float64,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineMemory(
		name string,
		megabytes int32,
	) (*virtual_machine.Machine, error)
	SetVirtualMachinePlatform(
		name string,
		platformName string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachinePrimaryAddress(
		name string,
		address string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineStatus(
		name string,
		status string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineTenant(
		name string,
		tenantName string,
	) (*virtual_machine.Machine, error)
	SiteByName(n string) (*site.Site, error)
	Sites() ([]*site.Site, error)
	Tags() ([]*tag.Tag, error)
	TenantByName(n string) (*tenant.Tenant, error)
	Tenants() ([]*tenant.Tenant, error)
	TunnelByName(n string) (*tunnel.Tunnel, error)
	TunnelGroupByName(n string) (*tunnel_group.Group, error)
	TunnelGroups() ([]*tunnel_group.Group, error)
	Tunnels() ([]*tunnel.Tunnel, error)
	TunnelTerminations() ([]*tunnel_termination.Termination, error)
	UpdateJournalEntry(
		identifier int32,
		kind string,
		comments string,
	) (*journal_entry.Entry, error)
	VirtualJournalEntries(
		machine string,
		limit int32,
		offset int32,
	) ([]*journal_entry.Entry, error)
	VirtualMachineAddresses(machine string) ([]*internet_address.Address, error)
	VirtualMachineByName(n string) (*virtual_machine.Machine, error)
	VirtualMachineDisks(m *virtual_machine.Machine) ([]*virtual_disk.Disk, error)
	VirtualMachineInterfaceByName(
		vm *virtual_machine.Machine,
		name string,
	) (*netbox.VMInterface, error)
	VirtualMachineInterfaces(m *virtual_machine.Machine) ([]*network.Interface, error)
	VirtualMachines() ([]*virtual_machine.Machine, error)
	WirelessNetworks() ([]*wireless_network.Network, error)
}
