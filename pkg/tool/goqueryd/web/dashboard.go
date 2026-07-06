package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	status, e := s.service.Status()

	if e != nil {
		s.view.RenderPage(
			w,
			constant.DashboardTitle,
			constant.DashboardPath,
			html.P(gomponents.Text("Failed to load status.")),
		)

		return
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.DashboardTitle)),
		html.Table(
			html.TBody(
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Documents"))),
					html.Td(gomponents.Textf("%d", status.TotalDocuments)),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Embeddings"))),
					html.Td(gomponents.Textf("%d", status.TotalEmbeddings)),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Pending"))),
					html.Td(gomponents.Textf("%d", status.PendingEmbeddings)),
				),
			),
		),
	)

	if len(status.Collections) > 0 {
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
					html.Td(
						html.Small(gomponents.Text(c.Path)),
					),
				),
			)
		}

		content = append(
			content,
			html.H3(gomponents.Text(constant.CollectionsTitle)),
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Name")),
						html.Th(gomponents.Text("Type")),
						html.Th(gomponents.Text("Documents")),
						html.Th(gomponents.Text("Path")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		content...,
	)
}
