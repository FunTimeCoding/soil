package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
	"github.com/funtimecoding/soil/pkg/netbox/contact"
	"github.com/funtimecoding/soil/pkg/netbox/contact_group"
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_link"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/module"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
	"github.com/funtimecoding/soil/pkg/netbox/module_type_profile"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
	"github.com/funtimecoding/soil/pkg/netbox/service"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/source"
	"github.com/funtimecoding/soil/pkg/netbox/system_number"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/netbox/user"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_link"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type Client interface {
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
	Bookmarks() ([]*bookmark.Bookmark, error)
	Cables() ([]*cable.Cable, error)
	ClusterByName(n string) (*cluster.Cluster, error)
	Clusters() ([]*cluster.Cluster, error)
	ClusterTypeByName(n string) (*cluster_type.Type, error)
	ClusterTypes() ([]*cluster_type.Type, error)
	ChoiceSetByName(n string) (*choice_set.Set, error)
	ChoiceSets() ([]*choice_set.Set, error)
	CreateChoiceSet(
		name string,
		choices []string,
	) (*choice_set.Set, error)
	CreateSelectionField(
		name string,
		label string,
		choiceSet string,
	) (*custom_field.Field, error)
	CustomFields() ([]*custom_field.Field, error)
	DeleteChoiceSet(name string) error
	DeleteCustomField(name string) error
	SetChoiceSetChoices(
		name string,
		choices []string,
	) (*choice_set.Set, error)
	SetDeviceCustomFields(
		name string,
		fields map[string]any,
	) (*device.Device, error)
	SetVirtualMachineCustomFields(
		name string,
		fields map[string]any,
	) (*virtual_machine.Machine, error)
	ConfigurationContextByName(
		n string,
	) (*configuration_context.Context, error)
	ConfigurationContexts() ([]*configuration_context.Context, error)
	CreateConfigurationContext(
		name string,
		weight int32,
		payload map[string]any,
		a *configuration_assignment.Assignment,
	) (*configuration_context.Context, error)
	DeleteConfigurationContext(name string) error
	SetConfigurationContextAssignment(
		name string,
		a *configuration_assignment.Assignment,
	) (*configuration_context.Context, error)
	SetConfigurationContextPayload(
		name string,
		payload map[string]any,
	) (*configuration_context.Context, error)
	SetConfigurationContextWeight(
		name string,
		weight int32,
	) (*configuration_context.Context, error)
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
	CreateDeviceRole(
		name string,
		virtualMachine bool,
	) (*device_role.Role, error)
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
	UpdateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
		h net.HardwareAddr,
	) (*network.Interface, error)
	PhysicalAddresses() ([]*physical_address.Address, error)
	UpdateVirtualInterface(
		vm *virtual_machine.Machine,
		name string,
		h net.HardwareAddr,
	) (*netbox.VMInterface, error)
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
	SetVirtualMachineRole(
		name string,
		roleName string,
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
	AssignPhysicalToInterface(
		p *physical_address.Address,
		i *network.Interface,
	) (*physical_address.Address, error)
	AssignPhysicalToVirtualInterface(
		p *physical_address.Address,
		identifier int32,
	) (*physical_address.Address, error)
	Close()
	ClusterGroups() ([]*cluster_group.Group, error)
	ClustersByName(s string) ([]*cluster.Cluster, error)
	ConfigurationTemplates() ([]*configuration_template.Template, error)
	ConsolePorts() ([]*console_port.Port, error)
	ConsoleServerPorts() ([]*console_server_port.Port, error)
	ContactGroups() ([]*contact_group.Group, error)
	ContactRoles() ([]*contact_role.Role, error)
	Contacts() ([]*contact.Contact, error)
	CreateInterfacePhysical(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
		h net.HardwareAddr,
	) (*network.Interface, error)
	CreateInternet(
		objectType string,
		objectIdentifier int64,
		i net.IP,
		m net.IPMask,
	) (*netbox.IPAddress, error)
	CreatePhysical(
		a net.HardwareAddr,
		description string,
	) (*physical_address.Address, error)
	CustomFieldByName(n string) (*custom_field.Field, error)
	CustomLinks() ([]*custom_link.Link, error)
	DeleteDevice(identifier int32) error
	DeleteInterface(identifier int32) error
	DeleteInterfaceByName(
		d *device.Device,
		name string,
	) error
	DeletePhysical(identifier int32) error
	DeviceBayTemplates() ([]*device_bay_template.Template, error)
	DeviceByNames(n []string) (*device.Device, error)
	DeviceCables(device string) ([]*cable.Cable, error)
	DeviceExists(name string) (bool, error)
	DeviceInterfaceExists(
		d *device.Device,
		name string,
	) (bool, error)
	DeviceModuleBays(device string) ([]*module_bay.Bay, error)
	DevicesByCluster(s string) ([]*device.Device, error)
	DevicesByName(s string) ([]*device.Device, error)
	DevicesByRack(i int32) ([]*device.Device, error)
	DevicesByTag(s string) ([]*device.Device, error)
	EnsurePhysical(
		h net.HardwareAddr,
	) (*physical_address.Address, error)
	ExportTemplates() ([]*export_template.Template, error)
	FindUnusedAddress(sub *net.IPNet) (net.IP, error)
	Forget()
	FrontPorts() ([]*front_port.Port, error)
	InterfaceAddresses(interfaceIdentifier int32) ([]*internet_address.Address, error)
	InternetAddresses() ([]*internet_address.Address, error)
	InventoryItemRoles() ([]*inventory_item_role.Role, error)
	InventoryItems() ([]*inventory_item.Item, error)
	LocationByName(n string) (*location.Location, error)
	ModuleBays() ([]*module_bay.Bay, error)
	ModuleTypeProfiles() ([]*module_type_profile.Profile, error)
	ModuleTypes() ([]*module_type.Type, error)
	Modules() ([]*module.Module, error)
	MustAddTag(
		deviceName string,
		tag string,
	) *device.Device
	MustBookmarks() []*bookmark.Bookmark
	MustCables() []*cable.Cable
	MustChoiceSets() []*choice_set.Set
	MustClusterGroups() []*cluster_group.Group
	MustClusterTypes() []*cluster_type.Type
	MustClusters() []*cluster.Cluster
	MustConfigurationContexts() []*configuration_context.Context
	MustConfigurationTemplates() []*configuration_template.Template
	MustConsolePorts() []*console_port.Port
	MustConsoleServerPorts() []*console_server_port.Port
	MustContactGroups() []*contact_group.Group
	MustContactRoles() []*contact_role.Role
	MustContacts() []*contact.Contact
	MustCreateDevice(
		name string,
		o *device_role.Role,
		tags []string,
		y *device_type.Type,
		s *site.Site,
		n *tenant.Tenant,
	) *device.Device
	MustCreateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
	) *network.Interface
	MustCreateInternet(
		objectType string,
		objectIdentifier int64,
		i net.IP,
		m net.IPMask,
	) *netbox.IPAddress
	MustCreatePhysical(
		a net.HardwareAddr,
		description string,
	) *physical_address.Address
	MustCustomFields() []*custom_field.Field
	MustCustomLinks() []*custom_link.Link
	MustDeleteInterface(identifier int32)
	MustDeviceAddresses(device string) []*internet_address.Address
	MustDeviceBayTemplates() []*device_bay_template.Template
	MustDeviceByName(n string) *device.Device
	MustDeviceCables(device string) []*cable.Cable
	MustDeviceInterfaceByName(
		d *device.Device,
		name string,
	) *network.Interface
	MustDeviceInterfaces(device int32) []*network.Interface
	MustDeviceModuleBays(device string) []*module_bay.Bay
	MustDeviceRoleByName(n string) *device_role.Role
	MustDeviceRoles() []*device_role.Role
	MustDeviceTagNames(name string) []string
	MustDeviceTypeByName(n string) *device_type.Type
	MustDeviceTypes() []*device_type.Type
	MustDevices() []*device.Device
	MustDevicesByMatch(s string) []*device.Device
	MustDevicesByName(s string) []*device.Device
	MustDevicesByRack(i int32) []*device.Device
	MustDevicesByTag(s string) []*device.Device
	MustExportTemplates() []*export_template.Template
	MustFindUnusedAddress(sub *net.IPNet) net.IP
	MustFrontPorts() []*front_port.Port
	MustInterfaceAddresses(interfaceIdentifier int32) []*internet_address.Address
	MustInternetAddressRanges() []*internet_address_range.Range
	MustInternetAddresses() []*internet_address.Address
	MustInventoryItemRoles() []*inventory_item_role.Role
	MustInventoryItems() []*inventory_item.Item
	MustLocations() []*location.Location
	MustManufacturers() []*manufacturer.Manufacturer
	MustModuleBays() []*module_bay.Bay
	MustModuleTypeProfiles() []*module_type_profile.Profile
	MustModuleTypes() []*module_type.Type
	MustModules() []*module.Module
	MustNotificationGroups() []*notification_group.Group
	MustPhysicalAddress(a net.HardwareAddr) *physical_address.Address
	MustPhysicalAddresses() []*physical_address.Address
	MustPhysicalAddressesByHardware(a net.HardwareAddr) []*physical_address.Address
	MustPlatforms() []*platform.Platform
	MustPowerFeeds() []*power_feed.Feed
	MustPowerOutlets() []*power_outlet.Outlet
	MustPowerPanels() []*power_panel.Panel
	MustPowerPorts() []*power_port.Port
	MustPrefixByDescription(d string) *prefix.Prefix
	MustPrefixByName(n string) *prefix.Prefix
	MustPrefixes() []*prefix.Prefix
	MustRackByName(n string) *rack.Rack
	MustRackRoles() []*rack_role.Role
	MustRackTypes() []*rack_type.Type
	MustRacks() []*rack.Rack
	MustRearPorts() []*rear_port.Port
	MustRemoveTag(
		deviceName string,
		tag string,
	) *device.Device
	MustServiceTemplates() []*service_template.Template
	MustServices() []*service.Service
	MustSiteByName(n string) *site.Site
	MustSites() []*site.Site
	MustSources() []*source.Source
	MustSystemNumbers() []*system_number.Number
	MustTags() []*tag.Tag
	MustTenantByName(n string) *tenant.Tenant
	MustTenantGroups() []*tenant_group.Group
	MustTenants() []*tenant.Tenant
	MustTunnelGroups() []*tunnel_group.Group
	MustTunnels() []*tunnel.Tunnel
	MustUpdateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
		h net.HardwareAddr,
	) *network.Interface
	MustUpdatePhysicalAddress(
		a *physical_address.Address,
		i *network.Interface,
	) *physical_address.Address
	MustUserGroups() []*user_group.Group
	MustUsers() []*user.User
	MustVirtualChassis() []*virtual_chassis.Chassis
	MustVirtualDeviceContexts() []*virtual_device_context.Context
	MustVirtualMachines() []*virtual_machine.Machine
	MustVirtualNetworkGroups() []*virtual_network_group.Group
	MustVirtualNetworks() []*virtual_network.Network
	MustWirelessLinks() []*wireless_link.Link
	MustWirelessNetworkGroups() []*wireless_network_group.Group
	MustWirelessNetworks() []*wireless_network.Network
	NotificationGroups() ([]*notification_group.Group, error)
	PhysicalAddress(a net.HardwareAddr) (*physical_address.Address, error)
	PhysicalAddressesByHardware(a net.HardwareAddr) ([]*physical_address.Address, error)
	PhysicalExists(a net.HardwareAddr) (bool, error)
	PlatformByName(n string) (*platform.Platform, error)
	PowerFeeds() ([]*power_feed.Feed, error)
	PowerOutlets() ([]*power_outlet.Outlet, error)
	PowerPanels() ([]*power_panel.Panel, error)
	PowerPorts() ([]*power_port.Port, error)
	PrefixByDescription(d string) (*prefix.Prefix, error)
	PrefixByName(n string) (*prefix.Prefix, error)
	RackByName(n string) (*rack.Rack, error)
	RackRoles() ([]*rack_role.Role, error)
	RackTypes() ([]*rack_type.Type, error)
	Racks() ([]*rack.Rack, error)
	RacksByName(n string) ([]*rack.Rack, error)
	Reachable(x context.Context) error
	Address() string
	RearPorts() ([]*rear_port.Port, error)
	Reservations() ([]*reservation.Reservation, error)
	ServiceTemplates() ([]*service_template.Template, error)
	Services() ([]*service.Service, error)
	SitesByMatch(m string) ([]*site.Site, error)
	SitesByName(n string) ([]*site.Site, error)
	Sources() ([]*source.Source, error)
	SystemNumbers() ([]*system_number.Number, error)
	TagByName(n string) (*tag.Tag, error)
	TenantGroups() ([]*tenant_group.Group, error)
	UnsetSerial(device string) (*device.Device, error)
	UpdatePhysicalAddress(
		a *physical_address.Address,
		i *network.Interface,
	) (*physical_address.Address, error)
	UpdateSerial(
		device string,
		serial string,
	) (*device.Device, error)
	UserGroups() ([]*user_group.Group, error)
	Users() ([]*user.User, error)
	VirtualChassis() ([]*virtual_chassis.Chassis, error)
	VirtualDeviceContexts() ([]*virtual_device_context.Context, error)
	VirtualNetworkGroups() ([]*virtual_network_group.Group, error)
	VirtualNetworks() ([]*virtual_network.Network, error)
	WirelessLinks() ([]*wireless_link.Link, error)
	WirelessNetworkGroups() ([]*wireless_network_group.Group, error)
}
