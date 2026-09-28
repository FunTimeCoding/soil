package argument

type UpdateClient struct {
	Identifier              string `json:"identifier"`
	RedirectLocator         string `json:"redirect_locator"`
	Scope                   string `json:"scope"`
	GrantType               string `json:"grant_type"`
	ResponseType            string `json:"response_type"`
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
}
