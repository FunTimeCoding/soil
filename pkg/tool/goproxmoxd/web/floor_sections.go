package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) floorSections(f floor.Floor) gomponents.Node {
	if len(f.Nodes) == 0 {
		return html.P(
			gomponents.Text("No hypervisors seen yet - first poll pending."),
		)
	}

	sections := []gomponents.Node{}

	for _, n := range f.Nodes {
		header := []gomponents.Node{
			html.Class("node-header"),
			html.H3(gomponents.Text(nodeLabel(n))),
			html.Span(
				html.Class(constant.ReleaseClass),
				gomponents.Text(n.Version),
			),
		}

		if n.UpdatesPending > 0 {
			header = append(
				header,
				html.Span(
					html.Class("badge updates"),
					gomponents.Textf("%d updates", n.UpdatesPending),
				),
			)
		}

		rows := []gomponents.Node{}

		for _, g := range f.Guests {
			if g.Hypervisor != n.Hypervisor || g.Node != n.Name {
				continue
			}

			rows = append(rows, s.guestRow(g))
		}

		sections = append(
			sections,
			html.Div(header...),
			html.Table(html.Class("floor-table"), html.TBody(rows...)),
		)
	}

	return html.Div(sections...)
}
