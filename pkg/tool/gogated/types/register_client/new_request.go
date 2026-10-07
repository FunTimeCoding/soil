package register_client

func NewRequest(
	redirectURIs []string,
	grantTypes []string,
	responseTypes []string,
	scopes []string,
	tokenEndpointAuthMethod string,
) *Request {
	return &Request{
		RedirectLocators:        redirectURIs,
		GrantTypes:              grantTypes,
		ResponseTypes:           responseTypes,
		Scopes:                  scopes,
		TokenEndpointAuthMethod: tokenEndpointAuthMethod,
	}
}
