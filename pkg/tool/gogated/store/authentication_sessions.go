package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"

func (s *Store) AuthenticationSessions() (
	[]*authentication_session.AuthenticationSession,
	error,
) {
	var result []*authentication_session.AuthenticationSession
	e := s.mapper.Order("last_used_at desc").Find(&result).Error

	return result, e
}
