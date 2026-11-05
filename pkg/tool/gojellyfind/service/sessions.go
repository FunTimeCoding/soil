package service

import "github.com/funtimecoding/soil/pkg/jellyfin/session"

func (s *Service) Sessions() ([]*session.Session, error) {
	return guarded(
		s.guard,
		func() ([]*session.Session, error) {
			return s.client.Sessions()
		},
	)
}
