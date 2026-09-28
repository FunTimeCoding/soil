package server

type registerResponse struct {
	ClientIdentifier        string   `json:"client_id"`
	ClientSecret            string   `json:"client_secret,omitempty"`
	RedirectLocators        []string `json:"redirect_uris"`
	GrantTypes              []string `json:"grant_types"`
	ResponseTypes           []string `json:"response_types"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method"`
}
