package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
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

		for i := len(impressions) - 1; i >= 0; i-- {
			imp := impressions[i]
			rows = append(
				rows,
				html.Tr(
					html.Td(
						html.Small(
							gomponents.Text(formatTime(imp.CreatedAt)),
						),
					),
					html.Td(
						html.Em(gomponents.Text(imp.Source)),
					),
					html.Td(gomponents.Text(imp.Content)),
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
