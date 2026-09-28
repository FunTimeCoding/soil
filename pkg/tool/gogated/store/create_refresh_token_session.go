package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/refresh_token"
	"github.com/ory/fosite"
)

func (s *Store) CreateRefreshTokenSession(
	_ context.Context,
	signature string,
	_ string,
	q fosite.Requester,
) error {
	clientIdentifier := q.GetClient().GetID()
	requestIdentifier, scopes, grantedScopes, session, form, requestedAt := toRequester(
		q,
		clientIdentifier,
	)

	return s.mapper.Create(
		refresh_token.New(
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
