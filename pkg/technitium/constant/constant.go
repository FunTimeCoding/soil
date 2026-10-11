package constant

const (
	HostEnvironment      = "TECHNITIUM_HOST"
	TokenEnvironment     = "TECHNITIUM_TOKEN"
	UntrustedEnvironment = "TECHNITIUM_UNTRUSTED"
)

const (
	Base       = "/api"
	OkayStatus = "ok"
)

const (
	SettingsGetPath    = "/settings/get"
	SettingsSetPath    = "/settings/set"
	ZoneOptionsGetPath = "/zones/options/get"
	ZoneOptionsSetPath = "/zones/options/set"
)

const (
	ZoneParameter                = "zone"
	AvailableKeyNamesParameter   = "includeAvailableTsigKeyNames"
	KeysParameter                = "tsigKeys"
	UpdateParameter              = "update"
	UpdateNetworkParameter       = "updateNetworkACL"
	UpdatePolicyParameter        = "updateSecurityPolicies"
	ZoneTransferParameter        = "zoneTransfer"
	ZoneTransferNetworkParameter = "zoneTransferNetworkACL"
	ZoneTransferKeyParameter     = "zoneTransferTsigKeyNames"
)

const (
	AccessDeny                 = "Deny"
	AccessAllow                = "Allow"
	AccessNameServer           = "AllowOnlyZoneNameServers"
	AccessNetwork              = "UseSpecifiedNetworkACL"
	AccessNameServerAndNetwork = "AllowZoneNameServersAndUseSpecifiedNetworkACL"
)

const (
	DefaultAlgorithm = "hmac-sha256"
	SecretSize       = 32
	ClearValue       = "false"
	RowSeparator     = "|"
	WildcardPrefix   = "*."
)
