package roster

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/roster_entry"
	"time"
)

func (s *Service) Roster(window time.Duration) ([]*roster_entry.Entry, error) {
	if window < 0 {
		return nil, validation.New(
			"seen_within must not be negative: %s",
			window,
		)
	}

	if window == 0 {
		window = constant.InactivityTimeoutWindow
	}

	now := s.clock()
	sessions, e := s.store.SeenSessions(now.Add(-window))

	if e != nil {
		return nil, e
	}

	present := now.Add(-constant.InactivityTimeoutWindow)
	result := make([]*roster_entry.Entry, 0, len(sessions))

	for _, v := range sessions {
		result = append(
			result,
			roster_entry.New(
				v.CallsignValue(),
				v.Identifier,
				v.LastSeen,
				v.LastSeen.After(present),
			),
		)
	}

	return result, nil
}
