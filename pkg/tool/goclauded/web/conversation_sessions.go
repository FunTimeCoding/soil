package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/enriched_session"
)

func (s *Server) conversationSessions(
	skip int,
) ([]*enriched_session.Session, int, error) {
	sessions, e := s.roster.EnrichedSessions(constant.SidebarLimit+1, skip)

	if e != nil {
		return nil, 0, e
	}

	if len(sessions) <= constant.SidebarLimit {
		return sessions, 0, nil
	}

	return sessions[:constant.SidebarLimit], skip + constant.SidebarLimit, nil
}
