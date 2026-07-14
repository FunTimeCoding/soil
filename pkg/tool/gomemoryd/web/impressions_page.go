package web

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) impressionsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	days := 7

	if v := r.URL.Query().Get("days"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			days = n
		}
	}

	since := time.Now().Add(
		-time.Duration(days) * 24 * time.Hour,
	).UTC().Format(time.RFC3339)
	impressions, e := s.service.RecentImpressions(since)

	if e != nil {
		s.view.RenderPage(
			w,
			constant.ImpressionsTitle,
			constant.ImpressionsPath,
			html.P(gomponents.Text("Failed to load impressions.")),
		)

		return
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.ImpressionsTitle)),
	)
	content = append(
		content,
		html.P(
			html.Small(
				impressionWindowLinks(days),
			),
		),
	)

	if len(impressions) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No impressions.")),
		)
	} else {
		var rows []gomponents.Node

		for _, i := range impressions {
			rows = append(
				rows,
				html.Tr(
					layout.TimeCell(
						library.Parse(time.RFC3339, i.CreatedAt),
					),
					html.Td(
						html.Em(gomponents.Text(i.Source)),
					),
					html.Td(gomponents.Text(i.Content)),
				),
			)
		}

		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Source")),
						html.Th(gomponents.Text("Content")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	s.view.RenderPage(
		w,
		constant.ImpressionsTitle,
		constant.ImpressionsPath,
		content...,
	)
}
