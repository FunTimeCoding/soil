package web

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) events(
	w http.ResponseWriter,
	r *http.Request,
) {
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	limit := 50
	offset := (page - 1) * limit
	o := store.NewQueryOption()
	o.Tool = r.URL.Query().Get(constant.Tool)
	o.Surface = r.URL.Query().Get(constant.Surface)
	o.Actor = r.URL.Query().Get(constant.Actor)
	o.Kind = r.URL.Query().Get(constant.Kind)
	o.Limit = limit + 1
	o.Offset = offset
	entries, queryError := s.store.Recent(o)
	errors.PanicOnError(queryError)
	hasMore := len(entries) > limit

	if hasMore {
		entries = entries[:limit]
	}

	var rows []gomponents.Node

	for _, e := range entries {
		cells := []gomponents.Node{
			layout.TimeCell(e.CreatedAt),
			html.Td(gomponents.Text(e.Tool)),
			html.Td(gomponents.Text(e.Surface)),
			html.Td(gomponents.Text(e.Actor)),
			html.Td(gomponents.Text(e.Kind)),
			html.Td(gomponents.Text(e.Outcome)),
		}

		if e.Detail != nil {
			var detail map[string]string

			if json.Unmarshal([]byte(*e.Detail), &detail) == nil {
				cells = append(cells, html.Td(detailNode(detail)))
			} else {
				cells = append(cells, html.Td())
			}
		} else {
			cells = append(cells, html.Td())
		}

		rows = append(rows, html.Tr(cells...))
	}

	var content []gomponents.Node
	content = append(content, html.H3(gomponents.Text(constant.EventsTitle)))
	content = append(content, filterBar(o))

	if len(rows) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No events recorded.")),
		)
	} else {
		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Tool")),
						html.Th(gomponents.Text("Surface")),
						html.Th(gomponents.Text("Actor")),
						html.Th(gomponents.Text("Kind")),
						html.Th(gomponents.Text("Outcome")),
						html.Th(gomponents.Text("Detail")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	var navigation []gomponents.Node

	if page > 1 {
		navigation = append(
			navigation,
			html.A(
				gomponents.Attr("href", eventsLink(o, page-1)),
				gomponents.Text("← Newer"),
			),
		)
	}

	if hasMore {
		if len(navigation) > 0 {
			navigation = append(navigation, gomponents.Text(" · "))
		}

		navigation = append(
			navigation,
			html.A(
				gomponents.Attr("href", eventsLink(o, page+1)),
				gomponents.Text("Older →"),
			),
		)
	}

	if len(navigation) > 0 {
		content = append(content, html.P(gomponents.Group(navigation)))
	}

	s.view.RenderPage(w, "", constant.EventsPath, content...)
}
