package roster

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/enriched_session"
)

func (s *Service) EnrichedSessions(
	limit int,
	offset int,
) ([]*enriched_session.Session, error) {
	sessions, e := s.store.AllSessions(limit, offset)

	if e != nil {
		return nil, e
	}

	cutoff := s.clock().Add(-constant.InactivityTimeoutWindow)
	result := make([]*enriched_session.Session, 0, len(sessions))

	for _, i := range sessions {
		result = append(
			result,
			enriched_session.FromStore(i, i.LastSeen.After(cutoff)),
		)
	}

	return result, nil
}
