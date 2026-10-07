package register_client

type Request struct {
	RedirectLocators        []string
	GrantTypes              []string
	ResponseTypes           []string
	Scopes                  []string
	TokenEndpointAuthMethod string
}
