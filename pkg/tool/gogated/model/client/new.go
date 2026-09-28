package client

import "time"

func New(
	identifier string,
	secret string,
	redirectURIs string,
	grantTypes string,
	responseTypes string,
	scopes string,
	public bool,
	tokenEndpointAuthMethod string,
) *Client {
	return &Client{
		Identifier:              identifier,
		Secret:                  secret,
		RedirectLocators:        redirectURIs,
		GrantTypes:              grantTypes,
		ResponseTypes:           responseTypes,
		Scopes:                  scopes,
		Public:                  public,
		TokenEndpointAuthMethod: tokenEndpointAuthMethod,
		CreatedAt:               time.Now(),
	}
}
