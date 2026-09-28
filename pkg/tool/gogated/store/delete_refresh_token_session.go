package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/refresh_token"
)

func (s *Store) DeleteRefreshTokenSession(
	_ context.Context,
	signature string,
) error {
	return s.mapper.Where("signature = ?", signature).
		Delete(refresh_token.Stub()).Error
}
