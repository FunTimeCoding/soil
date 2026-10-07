package register

func NewResponse(
	clientIdentifier string,
	clientSecret string,
	redirectLocators []string,
	grantTypes []string,
	responseTypes []string,
	tokenEndpointAuthMethod string,
) *Response {
	return &Response{
		ClientIdentifier:        clientIdentifier,
		ClientSecret:            clientSecret,
		RedirectLocators:        redirectLocators,
		GrantTypes:              grantTypes,
		ResponseTypes:           responseTypes,
		TokenEndpointAuthMethod: tokenEndpointAuthMethod,
	}
}
