package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/access_token"
)

func (s *Store) DeleteAccessTokenSession(
	_ context.Context,
	signature string,
) error {
	return s.mapper.Where("signature = ?", signature).
		Delete(access_token.Stub()).Error
}
