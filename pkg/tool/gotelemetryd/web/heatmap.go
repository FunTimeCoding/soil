package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) heatmap(
	w http.ResponseWriter,
	r *http.Request,
) {
	window := r.URL.Query().Get("window")

	if window == "" {
		window = "24h"
	}

	since := sinceFromWindow(window)
	groupBy := r.URL.Query().Get("group")

	if groupBy == "" {
		groupBy = "tool"
	}

	rows, queryError := s.store.Summary(since, "", groupBy)
	errors.PanicOnError(queryError)
	showSurface := groupBy == "surface"
	showKind := groupBy == "kind"
	var tableRows []gomponents.Node

	for _, row := range rows {
		cells := []gomponents.Node{
			html.Td(gomponents.Text(row.Tool)),
		}

		if showSurface {
			cells = append(cells, html.Td(gomponents.Text(row.Surface)))
		}

		if showKind {
			cells = append(cells, html.Td(gomponents.Text(row.Kind)))
		}

		cells = append(
			cells,
			html.Td(gomponents.Text(fmt.Sprintf("%d", row.Count))),
		)
		tableRows = append(tableRows, html.Tr(cells...))
	}

	headerCells := []gomponents.Node{
		html.Th(gomponents.Text("Tool")),
	}

	if showSurface {
		headerCells = append(headerCells, html.Th(gomponents.Text("Surface")))
	}

	if showKind {
		headerCells = append(headerCells, html.Th(gomponents.Text("Kind")))
	}

	headerCells = append(headerCells, html.Th(gomponents.Text("Count")))
	var content []gomponents.Node
	content = append(content, html.H3(gomponents.Text(constant.HeatmapTitle)))
	content = append(content, windowSelector(window, groupBy))

	if len(tableRows) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No events recorded.")),
		)
	} else {
		content = append(
			content,
			html.Table(
				html.THead(html.Tr(headerCells...)),
				html.TBody(tableRows...),
			),
		)
	}

	s.view.RenderPage(w, "", constant.HeatmapPath, content...)
}
