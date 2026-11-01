package entry_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	component "github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func EditRow(
	e *entry.Entry,
	alert string,
) gomponents.Node {
	mark := detailMark(e.Identifier)
	target := join.Empty("#", mark)

	return html.Tr(
		html.ID(mark),
		html.Class(constant.DetailRowClass),
		html.Td(
			gomponents.Attr("colspan", "6"),
			html.Div(
				html.Class(constant.DetailContentClass),
				component.New(fragment(constant.EditPath, e.Identifier), "Save").
					Swap(target, webConstant.SwapOuter).
					Alert(alert).
					Text(constant.Action, "Action", e.Action).
					Required().
					Text(constant.User, "User", e.User).
					Required().
					DateTime(
						constant.Timestamp,
						"Timestamp",
						e.Timestamp.Format(constant.FilterFormat),
					).
					Text(constant.System, "System", e.System).
					Text(constant.Service, "Service", e.Service).
					TextArea(constant.Description, "Description", e.Description).
					Lines("4").
					Beside(
						load.New(
							fragment(constant.DetailPath, e.Identifier),
							"Cancel",
						).
							Swap(target, webConstant.SwapOuter).
							Button().
							Class("outline secondary").
							Node(),
					).
					Node(),
			),
		),
	)
}
