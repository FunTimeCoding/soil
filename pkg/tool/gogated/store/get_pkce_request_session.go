package store

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/proof_key"
	"github.com/ory/fosite"
	"gorm.io/gorm"
)

func (s *Store) GetPKCERequestSession(
	_ context.Context,
	signature string,
	i fosite.Session,
) (fosite.Requester, error) {
	var row proof_key.Key
	e := s.mapper.Where("signature = ?", signature).First(&row).Error

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
		i,
		func(identifier string) (fosite.Client, error) {
			return s.getClientInternal(identifier)
		},
	)
}
