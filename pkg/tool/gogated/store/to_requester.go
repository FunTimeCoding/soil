package store

import (
	"github.com/ory/fosite"
	"time"
)

func toRequester(
	r fosite.Requester,
	clientIdentifier string,
) (string, string, string, string, string, time.Time) {
	return r.GetID(),
		marshalScopes(r.GetRequestedScopes()),
		marshalScopes(r.GetGrantedScopes()),
		marshalSession(r.GetSession()),
		marshalForm(r.GetRequestForm()),
		r.GetRequestedAt()
}
