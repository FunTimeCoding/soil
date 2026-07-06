package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) pulseSection(identifier string) gomponents.Node {
	pulses, e := s.service.PulsesBySession(identifier)
	errors.PanicOnError(e)
	var nodes []gomponents.Node

	if len(pulses) > 0 {
		nodes = append(
			nodes,
			html.H4(gomponents.Text("Pulses")),
		)
		var rows []gomponents.Node

		for _, p := range pulses {
			from := "→"

			if p.FromName != "" {
				from = fmt.Sprintf("%s →", p.FromName)
			}

			rows = append(
				rows,
				html.P(
					html.Small(
						gomponents.Textf(
							"%s ",
							p.CreatedAt.Format("15:04"),
						),
					),
					html.Em(gomponents.Textf("%s ", from)),
					gomponents.Text(p.Body),
				),
			)
		}

		nodes = append(nodes, gomponents.Group(rows))
	}

	nodes = append(
		nodes,
		html.Form(
			gomponents.Attr("method", "POST"),
			gomponents.Attr(
				"action",
				fmt.Sprintf("/sessions/%s/pulse", identifier),
			),
			html.Input(
				html.Type("text"),
				html.Name(constant.Body),
				gomponents.Attr("placeholder", "pulse..."),
				gomponents.Attr("autocomplete", "off"),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text("Send"),
			),
		),
	)

	return gomponents.Group(nodes)
}
