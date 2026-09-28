package proof_key

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
) *Key {
	return &Key{
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
