package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authorization_code"
	"github.com/ory/fosite"
)

func (s *Store) CreateAuthorizeCodeSession(
	_ context.Context,
	code string,
	q fosite.Requester,
) error {
	clientIdentifier := q.GetClient().GetID()
	requestIdentifier, scopes, grantedScopes, session, form, requestedAt := toRequester(
		q,
		clientIdentifier,
	)

	return s.mapper.Create(
		authorization_code.New(
			code,
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
