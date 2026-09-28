package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"

func (s *Service) AuthenticationSessions() (
	[]*authentication_session.AuthenticationSession,
	error,
) {
	return s.store.AuthenticationSessions()
}
