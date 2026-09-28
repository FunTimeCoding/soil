package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"
	"time"
)

func (s *Store) CleanExpiredLoginSessions() error {
	cutoff := time.Now().Add(-constant.LoginSessionTimeToLive)

	return s.mapper.Where("created_at < ?", cutoff).
		Delete(login_session.Stub()).Error
}
