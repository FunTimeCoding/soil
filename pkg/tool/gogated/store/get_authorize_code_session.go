package store

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authorization_code"
	"github.com/ory/fosite"
	"gorm.io/gorm"
)

func (s *Store) GetAuthorizeCodeSession(
	_ context.Context,
	code string,
	i fosite.Session,
) (fosite.Requester, error) {
	var row authorization_code.AuthorizationCode
	e := s.mapper.Where("signature = ?", code).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, fosite.ErrNotFound
		}

		return nil, e
	}

	r, e := fromRequester(
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

	if e != nil {
		return nil, e
	}

	if !row.Active {
		return r, fosite.ErrInvalidatedAuthorizeCode
	}

	return r, nil
}
