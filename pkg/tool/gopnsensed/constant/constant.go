package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gopnsensed",
	"OPNsense firewall bridge",
	"gopnsensed",
).WithInstructions(
	"OPNsense firewall - DHCP leases, host reservations, firewall rules, aliases, DNS forwards, blocklists, interfaces, firewall log and state table. Read-only.",
)

const (
	FirewallLog     = "firewall_log"
	FirewallStates  = "firewall_states"
	ListAliases     = "list_aliases"
	ListBlocklists  = "list_blocklists"
	ListForwards    = "list_forwards"
	ListHosts       = "list_hosts"
	ListInterfaces  = "list_interfaces"
	ListLeases      = "list_leases"
	ListPools       = "list_pools"
	ListRules       = "list_rules"
	ListSourceNat   = "list_source_nat"
)

const (
	HostEnvironment     = "GOPNSENSED_HOST"
	PortEnvironment     = "GOPNSENSED_PORT"
	InsecureEnvironment = "GOPNSENSED_INSECURE"
)

const DefaultLogLimit = 100
