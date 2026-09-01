package web

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) guestRow(g floor.Guest) gomponents.Node {
	dot := "status-dot"
	stopped := g.Status != constant.RunningStatus

	if !stopped {
		dot = join.Space(dot, "status-running")
	}

	cells := []gomponents.Node{
		html.Td(html.Span(html.Class(dot), html.Title(g.Status))),
		html.Td(gomponents.Textf("%d", g.Identifier)),
		html.Td(
			html.A(
				html.Href(s.guestLink(g)),
				html.Target("_blank"),
				gomponents.Text(g.Name),
			),
		),
		html.Td(html.Class("load"), gomponents.Text(load(g))),
	}

	if g.Unbacked {
		cells = append(
			cells,
			html.Td(
				html.Span(
					html.Class("badge unbacked"),
					gomponents.Text("unbacked"),
				),
			),
		)
	} else {
		cells = append(cells, html.Td())
	}

	if stopped {
		return html.Tr(html.Class("stopped"), gomponents.Group(cells))
	}

	return html.Tr(cells...)
}
