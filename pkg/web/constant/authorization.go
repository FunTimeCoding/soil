package constant

const (
	ProtectedResource = "/.well-known/oauth-protected-resource"
	ProtectedMethods  = "GET, OPTIONS"

	CallbackPortEnvironment = "CALLBACK_PORT"

	AuthorizationIssuerEnvironment           = "AUTHORIZATION_ISSUER"
	AuthorizationClientIdentifierEnvironment = "AUTHORIZATION_CLIENT_IDENTIFIER"
	AuthorizationClientSecretEnvironment     = "AUTHORIZATION_CLIENT_SECRET"
	AuthorizationEncryptionSecretEnvironment = "AUTHORIZATION_ENCRYPTION_SECRET"

	SignInPath   = "/sign-in"
	CallbackPath = "/callback"
	SignOutPath  = "/sign-out"
	SignOutTitle = "Sign out"

	AuthorizationFlowCookie    = "authorization_flow"
	AuthorizationSubjectCookie = "authorization_subject"
	AuthorizationDefaultScope  = "openid offline"
)

// Notation key
const (
	AuthorizationServer   = "authorization_servers"
	AuthorizationResource = "resource"
)
