package constant

const (
	ProtectedResource = "/.well-known/oauth-protected-resource"
	ProtectedMethods  = "GET, OPTIONS"

	PortEnvironment = "CALLBACK_PORT"

	ClientIdentifierEnvironment = "GATE_CLIENT_IDENTIFIER"
	ClientSecretEnvironment     = "GATE_CLIENT_SECRET"
	IssuerEnvironment           = "GATE_ISSUER"
	SecretEnvironment           = "GATE_SECRET"
	CallbackLocatorEnvironment  = "GATE_CALLBACK_LOCATOR"
	SignInPathEnvironment       = "GATE_SIGN_IN_PATH"

	FlowCookie    = "authorization_flow"
	SubjectCookie = "authorization_subject"
	DefaultScope  = "openid offline"
)

// Notation key
const (
	AuthorizationServer = "authorization_servers"
	Resource            = "resource"
)
