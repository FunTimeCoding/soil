package store

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/open_identity_session"
	"github.com/ory/fosite"
	"gorm.io/gorm"
)

func (s *Store) GetOpenIDConnectSession(
	_ context.Context,
	authorizeCode string,
	q fosite.Requester,
) (fosite.Requester, error) {
	var row open_identity_session.Session
	e := s.mapper.Where("signature = ?", authorizeCode).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, fosite.ErrNotFound
		}

		return nil, e
	}

	return fromRequester(
		row.RequestIdentifier,
		row.ClientIdentifier,
		row.Scopes,
		row.GrantedScopes,
		row.Session,
		row.Form,
		row.RequestedAt,
		q.GetSession(),
		func(identifier string) (fosite.Client, error) {
			return s.getClientInternal(identifier)
		},
	)
}
