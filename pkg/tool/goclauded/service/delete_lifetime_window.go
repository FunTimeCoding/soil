package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/receipt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) deleteLifetimeWindow(
	r *session.Session,
	result *receipt.Receipt,
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
