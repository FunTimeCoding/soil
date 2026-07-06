package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gonetboxd",
	"NetBox infrastructure inventory bridge",
	"gonetboxd",
).WithInstructions(
	"NetBox infrastructure inventory - devices, interfaces, IP addresses, sites, clusters, virtual machines. Both read and create operations.",
)

const (
	AddDeviceTag                   = "add_device_tag"
	AddVirtualTag                  = "add_virtual_tag"
	CreateAddress                  = "create_address"
	CreateCluster                  = "create_cluster"
	CreateClusterType              = "create_cluster_type"
	CreateDevice                   = "create_device"
	CreateDeviceRole               = "create_device_role"
	CreateDeviceTunnelTermination  = "create_device_tunnel_termination"
	CreateDeviceType               = "create_device_type"
	CreateInterface                = "create_interface"
	CreateManufacturer             = "create_manufacturer"
	CreatePrefix                   = "create_prefix"
	CreateSite                     = "create_site"
	CreateTag                      = "create_tag"
	CreateTenant                   = "create_tenant"
	CreateTunnel                   = "create_tunnel"
	CreateTunnelGroup              = "create_tunnel_group"
	CreateVirtualAddress           = "create_virtual_address"
	CreateVirtualInterface         = "create_virtual_interface"
	CreateVirtualMachine           = "create_virtual_machine"
	CreateVirtualTunnelTermination = "create_virtual_tunnel_termination"
	GetDevice                      = "get_device"
	ListAddresses                  = "list_addresses"
	ListClusterTypes               = "list_cluster_types"
	ListClusters                   = "list_clusters"
	ListDeviceRoles                = "list_device_roles"
	ListDeviceTypes                = "list_device_types"
	ListDevices                    = "list_devices"
	ListInterfaces                 = "list_interfaces"
	ListManufacturers              = "list_manufacturers"
	ListPrefixes                   = "list_prefixes"
	ListSites                      = "list_sites"
	ListTags                       = "list_tags"
	ListTenants                    = "list_tenants"
	ListTunnelGroups               = "list_tunnel_groups"
	ListTunnelTerminations         = "list_tunnel_terminations"
	ListTunnels                    = "list_tunnels"
	ListVirtualMachines            = "list_virtual_machines"
	RemoveDeviceTag                = "remove_device_tag"
	RemoveVirtualTag               = "remove_virtual_tag"

	Address        = "address"
	Cluster        = "cluster"
	ClusterType    = "cluster_type"
	Description    = "description"
	Device         = "device"
	Encapsulation  = "encapsulation"
	Group          = "group"
	Interface      = "interface"
	Manufacturer   = "manufacturer"
	Model          = "model"
	Prefix         = "prefix"
	Role           = "role"
	Site           = "site"
	Tag            = "tag"
	Tenant         = "tenant"
	Tunnel         = "tunnel"
	Type           = "type"
	VirtualMachine = "virtual_machine"
)
