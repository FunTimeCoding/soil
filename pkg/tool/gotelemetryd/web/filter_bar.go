package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func filterBar(o *query_option.Option) gomponents.Node {
	var active []gomponents.Node

	if o.Tool != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("tool=%s ", o.Tool)),
				html.A(
					gomponents.Attr("href", constant.EventsPath),
					gomponents.Text("×"),
				),
			),
		)
	}

	if o.Surface != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("surface=%s ", o.Surface)),
				html.A(
					gomponents.Attr("href", constant.EventsPath),
					gomponents.Text("×"),
				),
			),
		)
	}

	if o.Actor != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("actor=%s ", o.Actor)),
				html.A(
					gomponents.Attr("href", constant.EventsPath),
					gomponents.Text("×"),
				),
			),
		)
	}

	if o.Kind != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("kind=%s ", o.Kind)),
				html.A(
					gomponents.Attr("href", constant.EventsPath),
					gomponents.Text("×"),
				),
			),
		)
	}

	if len(active) == 0 {
		return gomponents.Text("")
	}

	return html.P(gomponents.Text("Filtered: "), gomponents.Group(active))
}
