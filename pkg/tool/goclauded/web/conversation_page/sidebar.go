package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/enriched_session"
	"maragu.dev/gomponents"
)

func Sidebar(
	sessions []*enriched_session.Session,
	current string,
	next int,
) gomponents.Node {
	var result []gomponents.Node

	for _, e := range sessions {
		result = append(result, sidebarEntry(e, e.Identifier == current))
	}

	if next > 0 {
		result = append(result, sentinel(next))
	}

	return gomponents.Group(result)
}
