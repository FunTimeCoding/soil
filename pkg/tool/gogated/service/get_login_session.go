package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"

func (s *Service) GetLoginSession(
	identifier string,
) (*login_session.LoginSession, error) {
	return s.store.GetLoginSession(identifier)
}
