package summary

import "time"

func New(
	identifier string,
	redirectLocator []string,
	scope []string,
	grantType []string,
	responseType []string,
	tokenEndpointAuthMethod string,
	public bool,
	createdAt time.Time,
) *Summary {
	return &Summary{
		Identifier:              identifier,
		RedirectLocator:         redirectLocator,
		Scope:                   scope,
		GrantType:               grantType,
		ResponseType:            responseType,
		TokenEndpointAuthMethod: tokenEndpointAuthMethod,
		Public:                  public,
		CreatedAt:               createdAt,
	}
}
