package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/session"

func (s *Service) Sessions() []*session.Session {
	return s.store.Sessions()
}
