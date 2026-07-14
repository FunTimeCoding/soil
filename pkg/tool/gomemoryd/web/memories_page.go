package web

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) memoriesPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	tag := r.URL.Query().Get(constant.Tag)
	memoryType := r.URL.Query().Get(constant.Type)
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	memories, e := s.service.ListMemories(memoryType, tag, true)

	if e != nil {
		s.view.RenderPage(
			w,
			constant.MemoriesTitle,
			constant.MemoriesPath,
			html.P(gomponents.Text("Failed to load memories.")),
		)

		return
	}

	limit := 20
	offset := (page - 1) * limit
	total := len(memories)

	if offset >= total {
		memories = nil
	} else {
		end := offset + limit

		if end > total {
			end = total
		}

		memories = memories[offset:end]
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.MemoriesTitle)),
	)

	if tag != "" || memoryType != "" {
		var filters []gomponents.Node

		if tag != "" {
			filters = append(
				filters,
				gomponents.Textf("tag: %s", tag),
			)
		}

		if memoryType != "" {
			if len(filters) > 0 {
				filters = append(filters, gomponents.Text(" · "))
			}

			filters = append(
				filters,
				gomponents.Textf("type: %s", memoryType),
			)
		}

		filters = append(
			filters,
			gomponents.Text(" · "),
			html.A(
				gomponents.Attr("href", constant.MemoriesPath),
				gomponents.Text("clear"),
			),
		)
		content = append(content, html.P(gomponents.Group(filters)))
	}

	if len(memories) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No memories.")),
		)
	} else {
		var rows []gomponents.Node

		for _, m := range memories {
			var pips []gomponents.Node

			for _, t := range m.Tags {
				pips = append(
					pips,
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf(
								"%s?%s=%s",
								constant.MemoriesPath,
								constant.Tag,
								t,
							),
						),
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
						html.Small(gomponents.Text(m.Description)),
					),
					html.Td(gomponents.Text(m.Type)),
					html.Td(gomponents.Group(pips)),
					layout.TimeCell(
						library.Parse(time.RFC3339, m.UpdatedAt),
					),
				),
			)
		}

		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Name")),
						html.Th(gomponents.Text("Description")),
						html.Th(gomponents.Text("Type")),
						html.Th(gomponents.Text("Tags")),
						html.Th(gomponents.Text("Updated")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	content = append(
		content,
		paginationLinks(
			page,
			offset,
			limit,
			total,
			tag,
			memoryType,
		)...,
	)
	s.view.RenderPage(
		w,
		constant.MemoriesTitle,
		constant.MemoriesPath,
		content...,
	)
}
