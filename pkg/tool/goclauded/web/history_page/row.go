package history_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func Row(
	e *timeline.Entry,
	editable bool,
) gomponents.Node {
	edit := editable && e.Identifier > 0 && isEditable(e.Kind)

	return html.Tr(
		gomponents.If(edit, html.ID(mark(e.Identifier))),
		layout.TimeCell(library.Parse(time.RFC3339, e.Timestamp)),
		html.Td(gomponents.Text(timeline.FormatDescription(e))),
		gomponents.If(
			editable,
			html.Td(
				gomponents.If(
					edit,
					load.New(EventPath(e.Identifier, true), "✎").
						Swap(target(e.Identifier), webConstant.SwapOuter).
						Node(),
				),
			),
		),
	)
}
