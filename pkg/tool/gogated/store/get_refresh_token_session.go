package store

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/refresh_token"
	"github.com/ory/fosite"
	"gorm.io/gorm"
)

func (s *Store) GetRefreshTokenSession(
	_ context.Context,
	signature string,
	i fosite.Session,
) (fosite.Requester, error) {
	var row refresh_token.RefreshToken
	e := s.mapper.Where("signature = ?", signature).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, fosite.ErrNotFound
		}

		return nil, e
	}

	if !row.Active {
		return nil, fosite.ErrInactiveToken
	}

	return fromRequester(
		row.RequestIdentifier,
		row.ClientIdentifier,
		row.Scopes,
		row.GrantedScopes,
		row.Session,
		row.Form,
		row.RequestedAt,
		i,
		func(identifier string) (fosite.Client, error) {
			return s.getClientInternal(identifier)
		},
	)
}
