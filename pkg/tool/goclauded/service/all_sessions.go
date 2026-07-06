package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Service) AllSessions(
	limit int,
	offset int,
) ([]session.Session, error) {
	return s.store.AllSessions(limit, offset)
}
