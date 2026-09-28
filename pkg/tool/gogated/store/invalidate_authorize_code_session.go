package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authorization_code"
)

func (s *Store) InvalidateAuthorizeCodeSession(
	_ context.Context,
	code string,
) error {
	return s.mapper.Model(authorization_code.Stub()).
		Where("signature = ?", code).
		Update("active", false).Error
}
