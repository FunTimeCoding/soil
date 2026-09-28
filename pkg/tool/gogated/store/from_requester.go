package store

import (
	"github.com/ory/fosite"
	"time"
)

func fromRequester(
	requestIdentifier string,
	clientIdentifier string,
	scopes string,
	grantedScopes string,
	session string,
	form string,
	requestedAt time.Time,
	s fosite.Session,
	clientGetter func(string) (fosite.Client, error),
) (fosite.Requester, error) {
	c, e := clientGetter(clientIdentifier)

	if e != nil {
		return nil, e
	}

	return &fosite.Request{
		ID:             requestIdentifier,
		RequestedAt:    requestedAt,
		Client:         c,
		RequestedScope: unmarshalScopes(scopes),
		GrantedScope:   unmarshalScopes(grantedScopes),
		Form:           unmarshalForm(form),
		Session:        unmarshalSession(session, s),
	}, nil
}
