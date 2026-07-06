package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) sessionDetailPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	d, e := s.service.SessionDetail(r.PathValue(constant.Identifier))
	errors.PanicOnError(e)

	if d == nil {
		s.view.RenderPage(
			w,
			"Session Not Found",
			constant.SessionsPath,
			html.P(gomponents.Text("Session not found.")),
		)

		return
	}

	title := d.Identifier[:8]

	if d.Alias != "" {
		title = d.Alias
	} else if d.Slug != "" {
		title = d.Slug
	}

	var content []gomponents.Node
	content = append(content, html.H3(gomponents.Text(title)))
	var identity []gomponents.Node

	if d.Session != nil && d.Session.Name != "" {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Name"))),
				html.Td(gomponents.Text(d.Session.Name)),
			),
		)
	}

	if d.Session != nil && d.Session.Callsign != nil {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Callsign"))),
				html.Td(gomponents.Text(*d.Session.Callsign)),
			),
		)
	}

	if d.Alias != "" {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Alias"))),
				html.Td(gomponents.Text(d.Alias)),
			),
		)
	}

	if d.Description != "" {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Description"))),
				html.Td(gomponents.Text(d.Description)),
			),
		)
	}

	if d.Slug != "" {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Slug"))),
				html.Td(html.Small(gomponents.Text(d.Slug))),
			),
		)
	}

	identity = append(
		identity,
		html.Tr(
			html.Td(html.Strong(gomponents.Text("Identifier"))),
			html.Td(html.Small(gomponents.Text(d.Identifier))),
		),
	)

	if d.Created != "" {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Created"))),
				html.Td(gomponents.Text(d.Created)),
			),
		)
	}

	if d.TurnCount > 0 {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Turns"))),
				html.Td(gomponents.Textf("%d", d.TurnCount)),
			),
		)
	}

	if d.Session != nil {
		identity = append(
			identity,
			html.Tr(
				html.Td(html.Strong(gomponents.Text("Last Seen"))),
				html.Td(gomponents.Text(relativeTime(d.Session.LastSeen))),
			),
		)
	}

	content = append(
		content,
		html.Table(html.TBody(identity...)),
	)
	labels, labelError := s.service.LabelsBySession(d.Identifier)
	errors.PanicOnError(labelError)

	if len(labels) > 0 {
		var pips []gomponents.Node

		for _, l := range labels {
			pips = append(
				pips,
				html.Span(
					html.Class("label-pip"),
					gomponents.Textf("(%s:%s)", l.Key, l.Value),
				),
			)
		}

		content = append(content, html.P(gomponents.Group(pips)))
	}

	if len(d.Completions) > 0 {
		content = append(
			content,
			html.H4(gomponents.Text("Completions")),
		)
		var rows []gomponents.Node

		for _, c := range d.Completions {
			rows = append(
				rows,
				html.Tr(
					html.Td(
						html.Small(
							gomponents.Text(
								c.CreatedAt.Format(time.DateMinute),
							),
						),
					),
					html.Td(gomponents.Textf("[%s]", c.Kind)),
					html.Td(gomponents.Text(c.Topic)),
					html.Td(html.Small(gomponents.Text(c.Summary))),
				),
			)
		}

		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Kind")),
						html.Th(gomponents.Text("Topic")),
						html.Th(gomponents.Text("Summary")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	if d.Summary != "" {
		content = append(
			content,
			html.H4(gomponents.Text("Summary")),
			html.Pre(gomponents.Text(d.Summary)),
		)
	}

	content = append(
		content,
		html.Div(
			gomponents.Attr("sse-swap", constant.Pulse),
			s.pulseSection(d.Identifier),
		),
	)
	content = append(
		content,
		html.P(
			html.A(
				gomponents.Attr(
					"href",
					fmt.Sprintf("/sessions/%s/edit", d.Identifier),
				),
				gomponents.Text("Edit"),
			),
			gomponents.Text(" · "),
			html.A(
				gomponents.Attr(
					"href",
					fmt.Sprintf("/conversations/%s", d.Identifier),
				),
				gomponents.Text("View conversation"),
			),
			gomponents.Text(" · "),
			html.A(
				gomponents.Attr(
					"hx-post",
					fmt.Sprintf("/sessions/%s/delete", d.Identifier),
				),
				gomponents.Attr(
					"hx-confirm",
					"Delete this session and all its data?",
				),
				html.Style("color: var(--pico-del-color); cursor: pointer"),
				gomponents.Text("Delete"),
			),
		),
	)
	s.view.RenderLivePage(
		w,
		title,
		constant.SessionsPath,
		fmt.Sprintf("subscribe=pulse&session=%s", d.Identifier),
		content...,
	)
}
