package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Service) FindSession(
	identifier string,
) (*session.Session, bool, error) {
	return s.store.FindSession(identifier)
}
