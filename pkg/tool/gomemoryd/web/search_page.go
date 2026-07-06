package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query().Get(constant.Query)
	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.SearchTitle)),
		html.Form(
			gomponents.Attr("method", "GET"),
			gomponents.Attr("action", constant.SearchPath),
			html.Input(
				html.Type("search"),
				html.Name(constant.Query),
				gomponents.Attr("placeholder", "search memories..."),
				gomponents.Attr("value", query),
				gomponents.Attr("autocomplete", "off"),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text(constant.SearchTitle),
			),
		),
	)

	if query != "" {
		results, e := s.service.SearchMemories(query, 20, "", "")

		if e != nil {
			content = append(
				content,
				html.P(gomponents.Textf("Search failed: %s", e.Error())),
			)
		} else if len(results) == 0 {
			content = append(
				content,
				html.P(gomponents.Text("No results.")),
			)
		} else {
			var rows []gomponents.Node

			for _, m := range results {
				var pips []gomponents.Node

				for _, t := range m.Tags {
					pips = append(
						pips,
						html.Span(
							html.Class("tag-pip"),
							gomponents.Text(t),
						),
						gomponents.Text(" "),
					)
				}

				rows = append(
					rows,
					html.Tr(
						html.Td(
							html.A(
								gomponents.Attr(
									"href",
									fmt.Sprintf(
										"/memories/%d",
										m.Identifier,
									),
								),
								gomponents.Text(m.Name),
							),
						),
						html.Td(
							html.Small(
								gomponents.Text(m.Description),
							),
						),
						html.Td(gomponents.Text(m.Type)),
						html.Td(gomponents.Group(pips)),
					),
				)
			}

			content = append(
				content,
				html.P(
					gomponents.Textf("%d results", len(results)),
				),
				html.Table(
					html.THead(
						html.Tr(
							html.Th(gomponents.Text("Name")),
							html.Th(gomponents.Text("Description")),
							html.Th(gomponents.Text("Type")),
							html.Th(gomponents.Text("Tags")),
						),
					),
					html.TBody(rows...),
				),
			)
		}
	}

	s.view.RenderPage(
		w,
		constant.SearchTitle,
		constant.SearchPath,
		content...,
	)
}
