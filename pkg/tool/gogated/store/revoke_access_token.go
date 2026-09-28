package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/access_token"
)

func (s *Store) RevokeAccessToken(
	_ context.Context,
	requestIdentifier string,
) error {
	return s.mapper.Model(access_token.Stub()).Where(
		"request_identifier = ?",
		requestIdentifier,
	).Update("active", false).Error
}
