package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) memoryDetailPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	raw := r.PathValue(constant.Identifier)
	identifier, e := strconv.ParseInt(raw, 10, 64)

	if e != nil {
		s.view.RenderPage(
			w,
			"Not Found",
			constant.MemoriesPath,
			html.P(gomponents.Text("Invalid memory identifier.")),
		)

		return
	}

	m, e := s.service.GetMemory(identifier)

	if e != nil {
		s.view.RenderPage(
			w,
			"Not Found",
			constant.MemoriesPath,
			html.P(gomponents.Text("Memory not found.")),
		)

		return
	}

	var content []gomponents.Node
	content = append(content, html.H3(gomponents.Text(m.Name)))
	content = append(
		content,
		html.Table(
			html.TBody(
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Description"))),
					html.Td(gomponents.Text(m.Description)),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Type"))),
					html.Td(gomponents.Text(m.Type)),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Created"))),
					html.Td(gomponents.Text(formatTime(m.CreatedAt))),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Updated"))),
					html.Td(gomponents.Text(formatTime(m.UpdatedAt))),
				),
				html.Tr(
					html.Td(html.Strong(gomponents.Text("Active"))),
					html.Td(gomponents.Textf("%v", m.IsActive)),
				),
			),
		),
	)

	if len(m.Tags) > 0 {
		var pips []gomponents.Node

		for _, t := range m.Tags {
			pips = append(
				pips,
				html.A(
					gomponents.Attr(
						"href",
						fmt.Sprintf("/memories?tag=%s", t),
					),
					html.Class("tag-pip"),
					gomponents.Text(t),
				),
				gomponents.Text(" "),
			)
		}

		content = append(content, html.P(gomponents.Group(pips)))
	}

	content = append(
		content,
		html.H4(gomponents.Text("Content")),
		html.Pre(gomponents.Text(m.Content)),
	)
	history, e := s.service.GetMemoryHistory(identifier)

	if e == nil && len(history) > 0 {
		var rows []gomponents.Node

		for _, v := range history {
			rows = append(
				rows,
				html.Tr(
					html.Td(
						html.Small(gomponents.Text(formatTime(v.ChangedAt))),
					),
					html.Td(gomponents.Text(v.ChangeType)),
					html.Td(html.Small(gomponents.Text(v.Source))),
					html.Td(html.Small(gomponents.Text(v.Description))),
				),
			)
		}

		content = append(
			content,
			html.H4(gomponents.Text("History")),
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Action")),
						html.Th(gomponents.Text("Source")),
						html.Th(gomponents.Text("Description")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	related, e := s.service.ListRelated(identifier)

	if e == nil && len(related) > 0 {
		var links []gomponents.Node

		for _, r := range related {
			links = append(
				links,
				html.Li(
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf("/memories/%d", r.Identifier),
						),
						gomponents.Text(r.Name),
					),
					gomponents.Textf(" - %s", r.Description),
				),
			)
		}

		content = append(
			content,
			html.H4(gomponents.Text("Relations")),
			html.Ul(links...),
		)
	}

	s.view.RenderPage(w, m.Name, constant.MemoriesPath, content...)
}
