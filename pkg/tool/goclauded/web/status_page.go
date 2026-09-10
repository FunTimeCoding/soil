package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) statusPage(
	w http.ResponseWriter,
	_ *http.Request,
) {
	findings, e := s.service.Findings()
	errors.PanicOnError(e)
	content := []gomponents.Node{html.H3(gomponents.Text(constant.StatusTitle))}

	if len(findings) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("Nothing inconsistent.")),
		)
		s.view.RenderPage(
			w,
			constant.StatusTitle,
			constant.StatusPath,
			content...,
		)

		return
	}

	var rows []gomponents.Node

	for _, i := range findings {
		rows = append(
			rows,
			html.Tr(
				html.Td(html.Code(gomponents.Text(i.Kind))),
				html.Td(statusSubject(i.Subject)),
				html.Td(gomponents.Text(i.Detail)),
			),
		)
	}

	content = append(
		content,
		html.Table(
			html.THead(
				html.Tr(
					html.Th(gomponents.Text("Kind")),
					html.Th(gomponents.Text("Subject")),
					html.Th(gomponents.Text("Detail")),
				),
			),
			html.TBody(rows...),
		),
	)
	s.view.RenderPage(
		w,
		constant.StatusTitle,
		constant.StatusPath,
		content...,
	)
}
