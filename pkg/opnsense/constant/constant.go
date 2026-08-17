package constant

const (
	HostEnvironment     = "OPNSENSE_HOST"
	KeyEnvironment      = "OPNSENSE_KEY"
	SecretEnvironment   = "OPNSENSE_SECRET"
	InsecureEnvironment = "OPNSENSE_INSECURE"

	Base = "api"
)

const (
	AliasSearch     = "firewall/alias/search_item"
	BlocklistSearch = "unbound/settings/search_dnsbl"
	ForwardSearch   = "unbound/settings/search_forward"
	HostSearch      = "dnsmasq/settings/search_host"
	InterfaceState  = "diagnostics/interface/get_interface_config"
	LeaseSearch     = "dnsmasq/leases/search"
	LogRead         = "diagnostics/firewall/log"
	PoolSearch      = "dnsmasq/settings/search_range"
	RuleSearch      = "firewall/filter/search_rule"
	SourceNatSearch = "firewall/source_nat/search_rule"
	StateQuery      = "diagnostics/firewall/query_states"
)
