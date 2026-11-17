package session

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/deletion_result"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) deleteLifetimeWindow(
	r *session.Session,
	result *deletion_result.Result,
) error {
	if r.Callsign == nil || *r.Callsign == "" {
		return nil
	}

	queued, e := s.store.DeleteQueueWindow(*r.Callsign, r.StartedAt, s.clock())

	if e != nil {
		return e
	}

	result.Queue = queued
	notified, f := s.store.DeleteNotificationWindow(
		*r.Callsign,
		r.StartedAt,
		s.clock(),
	)

	if f != nil {
		return f
	}

	result.Notifications = notified

	return nil
}
