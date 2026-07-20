package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) sessionsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	activeOnly := r.URL.Query().Get("active") == "true"
	all, e := s.service.EnrichedSessions(0, 0)
	errors.PanicOnError(e)
	var sessions []*enriched_session.Session

	if activeOnly {
		for _, e := range all {
			if e.Active {
				sessions = append(sessions, e)
			}
		}
	} else {
		sessions = all
	}

	limit := 20
	offset := (page - 1) * limit
	total := len(sessions)

	if offset >= total {
		sessions = nil
	} else {
		end := offset + limit

		if end > total {
			end = total
		}

		sessions = sessions[offset:end]
	}

	var rows []gomponents.Node

	for _, e := range sessions {
		rows = append(rows, sessionRow(e))
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.SessionsTitle)),
	)
	var filters []gomponents.Node

	if activeOnly {
		filters = append(
			filters,
			html.A(
				gomponents.Attr("href", constant.SessionsPath),
				gomponents.Text("All"),
			),
			gomponents.Text(" · "),
			html.Strong(gomponents.Text("Active")),
		)
	} else {
		filters = append(
			filters,
			html.Strong(gomponents.Text("All")),
			gomponents.Text(" · "),
			html.A(
				gomponents.Attr("href", "/sessions?active=true"),
				gomponents.Text("Active"),
			),
		)
	}

	content = append(
		content,
		html.P(html.Class("filter-toggle"), gomponents.Group(filters)),
	)

	if len(rows) == 0 {
		content = append(content, html.P(gomponents.Text("No sessions.")))
	} else {
		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Name")),
						html.Th(gomponents.Text("Description")),
						html.Th(gomponents.Text("Lines")),
						html.Th(gomponents.Text("Turns")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	var navigation []gomponents.Node
	base := constant.SessionsPath

	if activeOnly {
		base = "/sessions?active=true"
	}

	if page > 1 {
		navigation = append(
			navigation,
			html.A(
				gomponents.Attr(
					"href",
					sessionsPageLink(base, page-1),
				),
				gomponents.Text("← Newer"),
			),
		)
	}

	if offset+limit < total {
		if len(navigation) > 0 {
			navigation = append(navigation, gomponents.Text(" · "))
		}

		navigation = append(
			navigation,
			html.A(
				gomponents.Attr(
					"href",
					sessionsPageLink(base, page+1),
				),
				gomponents.Text("Older →"),
			),
		)
	}

	if len(navigation) > 0 {
		content = append(content, html.P(gomponents.Group(navigation)))
	}

	s.view.RenderPage(
		w,
		constant.SessionsTitle,
		constant.SessionsPath,
		content...,
	)
}
