package constant

import "regexp"

const (
	TenantEnvironment = "AZURE_TENANT_IDENTIFIER"
	ClientEnvironment = "AZURE_CLIENT_IDENTIFIER"
	GraphHost         = "graph.microsoft.com"
	LoginHost         = "login.microsoftonline.com"
	KeychainService   = "gographd"
)

const (
	DeviceCodePath = "/%s/oauth2/v2.0/devicecode"
	TokenPath      = "/%s/oauth2/v2.0/token"
)

const (
	ProfilePath  = "/v1.0/me"
	SelectFilter = "$select"
	TopFilter    = "$top"
	OrderFilter  = "$orderby"
	ExpandFilter = "$expand"
	MatchFilter  = "$filter"
	PageMaximum  = 50
)

const (
	ExpiredGrant          = "invalid_grant"
	InteractionRequired   = "interaction_required"
	InvalidClient         = "invalid_client"
	UnauthorizedClient    = "unauthorized_client"
	AccessDenied          = "access_denied"
	ExpiredDeviceCode     = "expired_token"
	DeclinedAuthorization = "authorization_declined"
)

const (
	SignInMissing       = "graph sign-in required: run gograph login"
	SignInExpired       = "graph sign-in expired: run gograph login"
	RegistrationRefused = "graph app registration refused the sign-in"
	SignInUnfinished    = "graph sign-in not finished"
)

var DirectoryCode = regexp.MustCompile(`AADSTS[0-9]+`)
var Scopes = []string{
	"User.Read",
	"offline_access",
	"Team.ReadBasic.All",
	"Channel.ReadBasic.All",
	"Chat.Read",
	"ChannelMessage.Read.All",
	"ChatMessage.Send",
	"Chat.Create",
	"User.ReadBasic.All",
}

const TeamsMessageKind = "message"
