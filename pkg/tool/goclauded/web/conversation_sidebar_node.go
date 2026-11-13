package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	"maragu.dev/gomponents"
	"strings"
)

func (s *Server) conversationSidebarNode(
	query string,
	kinds []string,
	current string,
) (gomponents.Node, error) {
	if strings.TrimSpace(query) == "" {
		sessions, next, e := s.conversationSessions(0)

		if e != nil {
			return nil, e
		}

		return conversation_page.Sidebar(sessions, current, next), nil
	}

	found, e := s.search.SearchConversations(query, kinds, 0)

	if e != nil {
		return nil, e
	}

	indexed, total := s.search.SearchProgress()

	return conversation_page.SearchResults(
		found,
		query,
		kinds,
		indexed,
		total,
	), nil
}
