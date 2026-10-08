package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"maragu.dev/gomponents"
)

func (s *Server) sidebarNodes(skip int) []gomponents.Node {
	sessions, e := s.service.EnrichedSessions(0, 0)
	errors.PanicOnError(e)
	limit := 30

	if skip >= len(sessions) {
		return nil
	}

	sessions = sessions[skip:]
	hasMore := len(sessions) > limit

	if hasMore {
		sessions = sessions[:limit]
	}

	var result []gomponents.Node

	for _, e := range sessions {
		result = append(result, entry(e))
	}

	if hasMore {
		result = append(result, sentinel(skip+limit))
	}

	return result
}
