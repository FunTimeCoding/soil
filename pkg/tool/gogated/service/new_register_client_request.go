package service

func NewRegisterClientRequest(
	redirectURIs []string,
	grantTypes []string,
	responseTypes []string,
	scopes []string,
	tokenEndpointAuthMethod string,
) *RegisterClientRequest {
	return &RegisterClientRequest{
		RedirectLocators:        redirectURIs,
		GrantTypes:              grantTypes,
		ResponseTypes:           responseTypes,
		Scopes:                  scopes,
		TokenEndpointAuthMethod: tokenEndpointAuthMethod,
	}
}
