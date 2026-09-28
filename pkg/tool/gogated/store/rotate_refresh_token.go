package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/refresh_token"
)

func (s *Store) RotateRefreshToken(
	_ context.Context,
	requestIdentifier string,
	_ string,
) error {
	return s.mapper.Model(refresh_token.Stub()).Where(
		"request_identifier = ?",
		requestIdentifier,
	).Update("active", false).Error
}
