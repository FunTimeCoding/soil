package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) statisticPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	summary, e := s.service.MemoryTokenSummary(
		r.URL.Query().Get(constant.Scope),
	)

	if e != nil {
		s.view.RenderPage(
			w,
			constant.StatisticTitle,
			constant.StatisticPath,
			html.P(gomponents.Text("Failed to compute token statistics.")),
		)

		return
	}

	content := []gomponents.Node{
		html.H3(gomponents.Text(constant.StatisticTitle)),
		html.P(gomponents.Text(constant.StatisticIntroduction)),
		spreadTable(summary),
	}
	var rows []gomponents.Node

	for _, c := range summary.Statistic {
		if c.Hidden {
			continue
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Textf("%d", c.Block)),
				html.Td(gomponents.Textf("%d", c.Description)),
				html.Td(memoryLink(c.Identifier, c.Name)),
				html.Td(tagText(c.Tags)),
			),
		)
	}

	content = append(
		content,
		html.Table(
			html.THead(
				html.Tr(
					html.Th(gomponents.Text("Block")),
					html.Th(gomponents.Text("Description")),
					html.Th(gomponents.Text("Memory")),
					html.Th(gomponents.Text("Tags")),
				),
			),
			html.TBody(rows...),
		),
	)

	if summary.Withheld > 0 {
		content = append(
			content,
			html.P(
				gomponents.Textf(
					"%d withheld from this listing; the spreads include them.",
					summary.Withheld,
				),
			),
		)
	}

	s.view.RenderPage(
		w,
		constant.StatisticTitle,
		constant.StatisticPath,
		content...,
	)
}
