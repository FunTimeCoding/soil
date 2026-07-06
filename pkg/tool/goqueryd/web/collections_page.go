package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) collectionsPage(
	w http.ResponseWriter,
	_ *http.Request,
) {
	status, e := s.service.Status()

	if e != nil {
		s.view.RenderPage(
			w,
			constant.CollectionsTitle,
			constant.CollectionsPath,
			html.P(gomponents.Text("Failed to load collections.")),
		)

		return
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.CollectionsTitle)),
	)

	if len(status.Collections) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No collections.")),
		)
	} else {
		var rows []gomponents.Node

		for _, c := range status.Collections {
			rows = append(
				rows,
				html.Tr(
					html.Td(
						html.A(
							gomponents.Attr(
								"href",
								fmt.Sprintf(
									"%s/%s",
									constant.CollectionsPath,
									c.Name,
								),
							),
							gomponents.Text(c.Name),
						),
					),
					html.Td(gomponents.Text(c.Type)),
					html.Td(gomponents.Textf("%d", c.DocumentCount)),
					html.Td(html.Small(gomponents.Text(c.Path))),
					html.Td(html.Small(gomponents.Text(c.Pattern))),
				),
			)
		}

		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Name")),
						html.Th(gomponents.Text("Type")),
						html.Th(gomponents.Text("Documents")),
						html.Th(gomponents.Text("Path")),
						html.Th(gomponents.Text("Pattern")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	s.view.RenderPage(
		w,
		constant.CollectionsTitle,
		constant.CollectionsPath,
		content...,
	)
}
