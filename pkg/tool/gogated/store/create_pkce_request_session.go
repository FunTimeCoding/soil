package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/proof_key"
	"github.com/ory/fosite"
)

func (s *Store) CreatePKCERequestSession(
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
		proof_key.New(
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
