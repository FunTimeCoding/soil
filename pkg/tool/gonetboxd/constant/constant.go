package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gonetboxd",
	"NetBox infrastructure inventory bridge",
	"gonetboxd",
).WithInstructions(
	"NetBox infrastructure inventory - devices, interfaces, IP addresses, sites, clusters, virtual machines. Both read and create operations.",
)

const (
	AddDeviceJournalEntry          = "add_device_journal_entry"
	AddDeviceTag                   = "add_device_tag"
	AddVirtualJournalEntry         = "add_virtual_journal_entry"
	AddVirtualTag                  = "add_virtual_tag"
	CreateAddress                  = "create_address"
	AnsibleInventory               = "ansible_inventory"
	CreateChoiceSet                = "create_choice_set"
	CreateCustomField              = "create_custom_field"
	DeleteChoiceSet                = "delete_choice_set"
	DeleteCustomField              = "delete_custom_field"
	ListChoiceSets                 = "list_choice_sets"
	ListCustomFields               = "list_custom_fields"
	SetDeviceCustomField           = "set_device_custom_field"
	SetVirtualCustomField          = "set_virtual_custom_field"
	UpdateChoiceSet                = "update_choice_set"
	CreateCluster                  = "create_cluster"
	CreateConfigurationContext     = "create_configuration_context"
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
	DeleteConfigurationContext     = "delete_configuration_context"
	DeleteJournalEntry             = "delete_journal_entry"
	GetConfigurationContext        = "get_configuration_context"
	GetDevice                      = "get_device"
	ListAddresses                  = "list_addresses"
	ListClusterTypes               = "list_cluster_types"
	ListClusters                   = "list_clusters"
	ListConfigurationContexts      = "list_configuration_contexts"
	ListDeviceRoles                = "list_device_roles"
	ListDeviceTypes                = "list_device_types"
	ListDevices                    = "list_devices"
	ListDeviceJournalEntries       = "list_device_journal_entries"
	ListDeviceLabels               = "list_device_labels"
	ListInterfaces                 = "list_interfaces"
	ListManufacturers              = "list_manufacturers"
	ListPrefixes                   = "list_prefixes"
	ListSites                      = "list_sites"
	ListTags                       = "list_tags"
	ListTenants                    = "list_tenants"
	ListTunnelGroups               = "list_tunnel_groups"
	ListTunnelTerminations         = "list_tunnel_terminations"
	ListTunnels                    = "list_tunnels"
	ListVirtualJournalEntries      = "list_virtual_journal_entries"
	ListVirtualLabels              = "list_virtual_labels"
	ListVirtualMachines            = "list_virtual_machines"
	RemoveDeviceLabel              = "remove_device_label"
	RemoveDeviceTag                = "remove_device_tag"
	RemoveVirtualLabel             = "remove_virtual_label"
	RemoveVirtualTag               = "remove_virtual_tag"
	SetDeviceLabel                 = "set_device_label"
	SetVirtualLabel                = "set_virtual_label"
	UpdateConfigurationContext     = "update_configuration_context"
	UpdateJournalEntry             = "update_journal_entry"

	Address         = "address"
	Cluster         = "cluster"
	ClusterType     = "cluster_type"
	Comments        = "comments"
	Description     = "description"
	Device          = "device"
	Encapsulation   = "encapsulation"
	Group           = "group"
	Interface       = "interface"
	Key             = "key"
	KeyColumn       = "key"
	Kind            = "kind"
	Manufacturer    = "manufacturer"
	Model           = "model"
	PhysicalAddress = "physical_address"
	Platform        = "platform"
	Prefix          = "prefix"
	Role            = "role"
	Site            = "site"
	Status          = "status"
	Tag             = "tag"
	Tags            = "tags"
	Tenant          = "tenant"
	Tunnel          = "tunnel"
	Type            = "type"
	Value           = "value"
	VirtualMachine  = "virtual_machine"

	VirtualMachineRole = "virtual_machine_role"

	Choices   = "choices"
	ChoiceSet = "choice_set"
	Field     = "field"
	Label     = "label"
	Values    = "values"
	Payload   = "payload"
	Weight    = "weight"
	Roles     = "roles"
	Clusters  = "clusters"
	Platforms = "platforms"
	Sites     = "sites"
	Tenants   = "tenants"
)

const ConfigurationContextWeight int32 = 1000
const (
	AnsibleField = "ansible_playbook"
	SaltField    = "salt_state"
)
const (
	BookmarkTitle = "Bookmarks"
	BookmarkPath  = "/"
	BookmarkNone  = "Nothing bookmarked yet."
)

const (
	HostEnvironment     = "GONETBOX_HOST"
	TokenEnvironment    = "GONETBOX_TOKEN"
	PortEnvironment     = "GONETBOX_PORT"
	InsecureEnvironment = "GONETBOX_INSECURE"

	PowerIntentEnvironment = "GONETBOX_POWER_INTENT"
)

const (
	IntentInterval = time.Minute

	StatusMetric = "netbox_status"
	StatusHelp   = "A device or virtual machine's netbox status, 1 for the status it holds"
	KindLabel    = "kind"
	NameLabel    = "name"
	StatusLabel  = "status"
	DeviceKind   = "device"
	MachineKind  = "virtual_machine"
)

const ScopeNameKey = "name"
const (
	StatusActive  = "active"
	StatusOffline = "offline"
)
