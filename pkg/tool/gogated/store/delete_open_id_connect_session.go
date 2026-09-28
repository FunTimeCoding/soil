package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/open_identity_session"
)

func (s *Store) DeleteOpenIDConnectSession(
	_ context.Context,
	authorizeCode string,
) error {
	return s.mapper.Where("signature = ?", authorizeCode).
		Delete(open_identity_session.Stub()).Error
}
