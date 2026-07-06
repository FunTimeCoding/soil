package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) documentDetailPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	path := r.PathValue(constant.Path)
	document, suggestions, e := s.service.GetDocument(path)

	if e != nil {
		s.view.RenderPage(
			w,
			"Error",
			constant.SearchPath,
			html.P(
				gomponents.Textf(
					"Failed to load document: %s",
					e.Error(),
				),
			),
		)

		return
	}

	if document == nil {
		var content []gomponents.Node
		content = append(
			content,
			html.H3(gomponents.Text("Not Found")),
			html.P(gomponents.Textf("No document at: %s", path)),
		)

		if len(suggestions) > 0 {
			content = append(
				content,
				html.P(gomponents.Text("Did you mean:")),
			)
			var items []gomponents.Node

			for _, suggestion := range suggestions {
				items = append(
					items,
					html.Li(gomponents.Text(suggestion)),
				)
			}

			content = append(content, html.Ul(items...))
		}

		s.view.RenderPage(
			w,
			"Not Found",
			constant.SearchPath,
			content...,
		)

		return
	}

	var content []gomponents.Node
	content = append(content, html.H3(gomponents.Text(document.Title)))
	content = append(
		content,
		html.Table(
			html.TBody(
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Collection"))),
					html.Td(gomponents.Text(document.Collection)),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Path"))),
					html.Td(gomponents.Text(document.Path)),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Virtual Path"))),
					html.Td(
						html.Small(gomponents.Text(document.VirtualPath)),
					),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("File"))),
					html.Td(
						html.Small(gomponents.Text(document.FilePath)),
					),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Hash"))),
					html.Td(
						html.Small(gomponents.Text(document.Hash[:16])),
					),
				),
			),
		),
	)
	content = append(
		content,
		html.H4(gomponents.Text("Body")),
		html.Pre(gomponents.Text(document.Body)),
	)
	s.view.RenderPage(
		w,
		document.Title,
		constant.CollectionsPath,
		content...,
	)
}
