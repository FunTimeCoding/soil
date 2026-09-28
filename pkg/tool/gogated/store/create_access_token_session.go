package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/access_token"
	"github.com/ory/fosite"
)

func (s *Store) CreateAccessTokenSession(
	_ context.Context,
	signature string,
	q fosite.Requester,
) error {
	clientIdentifier := q.GetClient().GetID()
	requestIdentifier, scopes, grantedScopes, session, form, requestedAt := toRequester(
		q,
		clientIdentifier,
	)

	return s.mapper.Create(
		access_token.New(
			signature,
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
