package access_token

import "time"

func New(
	signature string,
	requestIdentifier string,
	clientIdentifier string,
	scopes string,
	grantedScopes string,
	session string,
	form string,
	requestedAt time.Time,
) *AccessToken {
	return &AccessToken{
		Signature:         signature,
		RequestIdentifier: requestIdentifier,
		ClientIdentifier:  clientIdentifier,
		Scopes:            scopes,
		GrantedScopes:     grantedScopes,
		Session:           session,
		Form:              form,
		Active:            true,
		RequestedAt:       requestedAt,
	}
}
