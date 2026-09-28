package service

type RegisterClientRequest struct {
	RedirectLocators        []string
	GrantTypes              []string
	ResponseTypes           []string
	Scopes                  []string
	TokenEndpointAuthMethod string
}
