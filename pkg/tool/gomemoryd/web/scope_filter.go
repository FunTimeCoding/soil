package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) scopeFilter(
	current string,
	tag string,
	memoryType string,
) []gomponents.Node {
	scopes, e := s.service.ListScopes()

	if e != nil || len(scopes) < 2 {
		return nil
	}

	entries := []gomponents.Node{
		html.Small(gomponents.Text("Scope: ")),
	}
	entries = append(
		entries,
		scopeLink(constant.AllScope, "", current, tag, memoryType),
	)

	for _, scope := range scopes {
		label := scope

		if scope == "" {
			label = constant.DefaultScope
		}

		entries = append(
			entries,
			gomponents.Text(" · "),
			scopeLink(label, label, current, tag, memoryType),
		)
	}

	return []gomponents.Node{html.P(entries...)}
}
