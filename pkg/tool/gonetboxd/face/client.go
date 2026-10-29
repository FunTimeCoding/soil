package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type Client interface {
	AddDeviceJournalEntry(
		x context.Context,
		device string,
		kind string,
		comments string,
	) (client.JournalEntry, error)
	AddDeviceTag(
		x context.Context,
		device string,
		tag string,
	) (client.Device, error)
	AddVirtualJournalEntry(
		x context.Context,
		machine string,
		kind string,
		comments string,
	) (client.JournalEntry, error)
	AddVirtualTag(
		x context.Context,
		name string,
		tag string,
	) (client.VirtualMachine, error)
	Address() string
	AnsibleInventoryMarkup(
		x context.Context,
		field string,
	) (string, error)
	Close()
	CreateAddress(
		x context.Context,
		device string,
		interfaceName string,
		address string,
		status string,
	) (client.Address, error)
	CreateAddressRange(
		x context.Context,
		start string,
		end string,
		status string,
		description string,
	) (client.AddressRange, error)
	CreateCable(
		x context.Context,
		deviceA string,
		interfaceA string,
		deviceB string,
		interfaceB string,
	) (client.Cable, error)
	CreateChoiceSet(
		x context.Context,
		name string,
		choices []string,
	) (client.ChoiceSet, error)
	CreateCluster(
		x context.Context,
		name string,
		clusterType string,
		site string,
	) (client.Cluster, error)
	CreateClusterType(
		x context.Context,
		name string,
	) (client.ClusterType, error)
	CreateConfigurationContext(
		x context.Context,
		body client.CreateConfigurationContextRequest,
	) (client.ConfigurationContext, error)
	CreateCustomField(
		x context.Context,
		name string,
		label string,
		choiceSet string,
	) (client.CustomField, error)
	CreateDevice(
		x context.Context,
		name string,
		role string,
		deviceType string,
		site string,
		tenant *string,
	) (client.Device, error)
	CreateDeviceRole(
		x context.Context,
		name string,
		virtualMachine bool,
	) (client.DeviceRole, error)
	CreateDeviceTunnelTermination(
		x context.Context,
		device string,
		tunnel string,
		interfaceName string,
		role string,
	) (client.TunnelTermination, error)
	CreateDeviceType(
		x context.Context,
		model string,
		manufacturer string,
	) (client.DeviceType, error)
	CreateInterface(
		x context.Context,
		device string,
		name string,
		interfaceType string,
		physicalAddress string,
	) (client.Interface, error)
	CreateLocation(
		x context.Context,
		name string,
		site string,
	) (client.Location, error)
	CreateManufacturer(
		x context.Context,
		name string,
	) (client.Manufacturer, error)
	CreatePlatform(
		x context.Context,
		name string,
	) (client.Platform, error)
	CreatePrefix(
		x context.Context,
		prefix string,
		site string,
		description string,
	) (client.Prefix, error)
	CreateSite(
		x context.Context,
		name string,
	) (client.Site, error)
	CreateTag(
		x context.Context,
		name string,
	) (client.Tag, error)
	CreateTenant(
		x context.Context,
		name string,
	) (client.Tenant, error)
	CreateTunnel(
		x context.Context,
		name string,
		encapsulation string,
		group string,
	) (client.Tunnel, error)
	CreateTunnelGroup(
		x context.Context,
		name string,
	) (client.TunnelGroup, error)
	CreateVirtualAddress(
		x context.Context,
		vmName string,
		interfaceName string,
		address string,
		status string,
	) (client.Address, error)
	CreateVirtualDisk(
		x context.Context,
		machine string,
		name string,
		size int32,
	) (client.VirtualDisk, error)
	CreateVirtualInterface(
		x context.Context,
		vmName string,
		name string,
	) (client.VirtualInterface, error)
	CreateVirtualMachine(
		x context.Context,
		name string,
		cluster string,
	) (client.VirtualMachine, error)
	CreateVirtualTunnelTermination(
		x context.Context,
		vm string,
		tunnel string,
		interfaceName string,
		role string,
	) (client.TunnelTermination, error)
	CreateWirelessNetwork(
		x context.Context,
		ssid string,
	) (client.WirelessNetwork, error)
	DeleteAddress(
		x context.Context,
		identifier int32,
	) error
	DeleteChoiceSet(
		x context.Context,
		name string,
	) error
	DeleteConfigurationContext(
		x context.Context,
		name string,
	) error
	DeleteCustomField(
		x context.Context,
		name string,
	) error
	DeleteJournalEntry(
		x context.Context,
		identifier int32,
	) error
	GetChoiceSet(
		x context.Context,
		name string,
	) (client.ChoiceSet, error)
	GetConfigurationContext(
		x context.Context,
		name string,
	) (client.ConfigurationContext, error)
	GetDevice(
		x context.Context,
		name string,
	) (client.Device, error)
	GetVirtualMachine(
		x context.Context,
		name string,
	) (client.VirtualMachine, error)
	ListAddressRanges(
		x context.Context,
	) ([]*client.AddressRange, error)
	ListAddresses(
		x context.Context,
		name string,
	) ([]*client.Address, error)
	ListCables(
		x context.Context,
	) ([]*client.Cable, error)
	ListChoiceSets(
		x context.Context,
	) ([]*client.ChoiceSet, error)
	ListClusterTypes(
		x context.Context,
	) ([]*client.ClusterType, error)
	ListClusters(
		x context.Context,
	) ([]*client.Cluster, error)
	ListConfigurationContexts(
		x context.Context,
	) ([]*client.ConfigurationContext, error)
	ListCustomFields(
		x context.Context,
	) ([]*client.CustomField, error)
	ListDeviceJournalEntries(
		x context.Context,
		device string,
		limit int32,
		offset int32,
	) ([]*client.JournalEntry, error)
	ListDeviceLabels(
		x context.Context,
		device string,
	) ([]*client.Label, error)
	ListDeviceTags(
		x context.Context,
		device string,
	) ([]string, error)
	ListDeviceRoles(
		x context.Context,
	) ([]*client.DeviceRole, error)
	ListDeviceTypes(
		x context.Context,
	) ([]*client.DeviceType, error)
	ListDevices(
		x context.Context,
		query *string,
	) ([]*client.Device, error)
	ListInterfaces(
		x context.Context,
		name string,
	) ([]*client.Interface, error)
	ListLocations(
		x context.Context,
	) ([]*client.Location, error)
	ListManufacturers(
		x context.Context,
	) ([]*client.Manufacturer, error)
	ListPhysicalAddresses(
		x context.Context,
	) ([]client.PhysicalAddressOwner, error)
	ListPlatforms(
		x context.Context,
	) ([]*client.Platform, error)
	ListPrefixes(
		x context.Context,
	) ([]*client.Prefix, error)
	ListSites(
		x context.Context,
	) ([]*client.Site, error)
	ListTags(
		x context.Context,
	) ([]*client.Tag, error)
	ListTenants(
		x context.Context,
	) ([]*client.Tenant, error)
	ListTunnelGroups(
		x context.Context,
	) ([]*client.TunnelGroup, error)
	ListTunnelTerminations(
		x context.Context,
	) ([]*client.TunnelTermination, error)
	ListTunnels(
		x context.Context,
	) ([]*client.Tunnel, error)
	ListVirtualAddresses(
		x context.Context,
		name string,
	) ([]*client.Address, error)
	ListVirtualDisks(
		x context.Context,
		machine string,
	) ([]*client.VirtualDisk, error)
	ListVirtualInterfaces(
		x context.Context,
		name string,
	) ([]*client.Interface, error)
	ListVirtualJournalEntries(
		x context.Context,
		machine string,
		limit int32,
		offset int32,
	) ([]*client.JournalEntry, error)
	ListVirtualLabels(
		x context.Context,
		machine string,
	) ([]*client.Label, error)
	ListVirtualMachines(
		x context.Context,
	) ([]*client.VirtualMachine, error)
	ListWirelessNetworks(
		x context.Context,
	) ([]*client.WirelessNetwork, error)
	Reachable(x context.Context) error
	RemoveDeviceLabel(
		x context.Context,
		device string,
		key string,
	) error
	RemoveDeviceTag(
		x context.Context,
		device string,
		tag string,
	) (client.Device, error)
	RemoveVirtualLabel(
		x context.Context,
		machine string,
		key string,
	) error
	RemoveVirtualTag(
		x context.Context,
		name string,
		tag string,
	) (client.VirtualMachine, error)
	SetDeviceLabel(
		x context.Context,
		device string,
		key string,
		value string,
	) (client.Label, error)
	SetInterfacePhysicalAddress(
		x context.Context,
		device string,
		name string,
		address string,
	) (client.PhysicalAddress, error)
	SetVirtualInterfacePhysicalAddress(
		x context.Context,
		machine string,
		name string,
		address string,
	) (client.PhysicalAddress, error)
	SetVirtualLabel(
		x context.Context,
		machine string,
		key string,
		value string,
	) (client.Label, error)
	UpdateChoiceSet(
		x context.Context,
		name string,
		choices []string,
	) (client.ChoiceSet, error)
	UpdateConfigurationContext(
		x context.Context,
		name string,
		body client.UpdateConfigurationContextRequest,
	) (client.ConfigurationContext, error)
	UpdateDevice(
		x context.Context,
		name string,
		body client.UpdateDeviceRequest,
	) (client.Device, error)
	UpdateJournalEntry(
		x context.Context,
		identifier int32,
		kind string,
		comments string,
	) (client.JournalEntry, error)
	UpdateVirtualMachine(
		x context.Context,
		name string,
		body client.UpdateVirtualMachineRequest,
	) (client.VirtualMachine, error)
}
