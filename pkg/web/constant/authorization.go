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

	AuthorizationFlowCookie    = "authorization_flow"
	AuthorizationSubjectCookie = "authorization_subject"
	AuthorizationDefaultScope  = "openid offline"
	// Notation key
	AuthorizationServer   = "authorization_servers"
	AuthorizationResource = "resource"
)
