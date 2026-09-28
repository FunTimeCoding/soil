package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/open_identity_session"
	"github.com/ory/fosite"
)

func (s *Store) CreateOpenIDConnectSession(
	_ context.Context,
	authorizeCode string,
	q fosite.Requester,
) error {
	clientIdentifier := q.GetClient().GetID()
	requestIdentifier, scopes, grantedScopes, session, form, requestedAt := toRequester(
		q,
		clientIdentifier,
	)

	return s.mapper.Create(
		open_identity_session.New(
			authorizeCode,
			requestIdentifier,
			clientIdentifier,
			scopes,
			grantedScopes,
			session,
			form,
			requestedAt,
		),
	).Error
}
