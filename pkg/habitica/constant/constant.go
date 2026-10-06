package constant

const (
	HostEnvironment  = "HABITICA_HOST"
	UserEnvironment  = "HABITICA_USER"
	TokenEnvironment = "HABITICA_TOKEN" // #nosec G101 not a hardcoded secret

	Base = "/api/v3"

	UserHeader  = "x-api-user"
	TokenHeader = "x-api-key"

	StatParameter = "stat"
	TypeParameter = "type"
)
