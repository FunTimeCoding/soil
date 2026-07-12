package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
	"time"
)

func (s *Service) UnnamedSessions(
	limit int,
	offset int,
) ([]*enriched_session.Session, error) {
	sessions, e := s.store.UnnamedSessions(limit, offset)

	if e != nil {
		return nil, e
	}

	cutoff := s.clock().Add(-1 * time.Hour)
	result := make([]*enriched_session.Session, 0, len(sessions))

	for _, i := range sessions {
		result = append(
			result,
			enriched_session.FromStore(i, i.LastSeen.After(cutoff)),
		)
	}

	return result, nil
}
