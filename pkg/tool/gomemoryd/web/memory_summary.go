package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) memorySummary() gomponents.Node {
	memories, e := s.service.ListMemories("", "", true)

	if e != nil {
		return html.P(gomponents.Text("Failed to load memories."))
	}

	counts := map[string]int{}

	for _, m := range memories {
		counts[m.Type]++
	}

	var rows []gomponents.Node

	for _, t := range []string{"feedback", "user", "project", "reference"} {
		if c, okay := counts[t]; okay {
			rows = append(
				rows,
				html.Tr(
					html.Td(gomponents.Text(t)),
					html.Td(gomponents.Textf("%d", c)),
				),
			)
		}
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.MemoriesTitle)),
			html.P(gomponents.Textf("%d active", len(memories))),
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Type")),
						html.Th(gomponents.Text("Count")),
					),
				),
				html.TBody(rows...),
			),
		},
	)
}
