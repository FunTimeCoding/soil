package constant

const (
	ProtectedResource = "/.well-known/oauth-protected-resource"
	ProtectedMethods  = "GET, OPTIONS"

	CallbackPortEnvironment = "CALLBACK_PORT"

	GateClientIdentifierEnvironment = "GATE_CLIENT_IDENTIFIER"
	GateClientSecretEnvironment     = "GATE_CLIENT_SECRET"
	GateIssuerEnvironment           = "GATE_ISSUER"
	GateSecretEnvironment           = "GATE_SECRET"
	GateCallbackLocatorEnvironment  = "GATE_CALLBACK_LOCATOR"
	GateSignInPathEnvironment       = "GATE_SIGN_IN_PATH"

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
	// Notation key
	AuthorizationServer   = "authorization_servers"
	AuthorizationResource = "resource"
)
