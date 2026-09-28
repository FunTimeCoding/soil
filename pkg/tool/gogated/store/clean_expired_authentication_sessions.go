package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"time"
)

func (s *Store) CleanExpiredAuthenticationSessions() error {
	now := time.Now()

	return s.mapper.Where(
		"last_used_at < ? OR created_at < ?",
		now.Add(-constant.AuthenticationIdleTimeToLive),
		now.Add(-constant.AuthenticationAbsoluteTimeToLive),
	).Delete(authentication_session.Stub()).Error
}
