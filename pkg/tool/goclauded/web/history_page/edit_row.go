package history_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func EditRow(e *event.Event) gomponents.Node {
	return html.Tr(
		html.ID(mark(e.Identifier)),
		layout.TimeCell(e.CreatedAt),
		html.Td(
			gomponents.Attr("colspan", "2"),
			entry.New(EventPath(e.Identifier, true), "Save").
				Swap(target(e.Identifier), webConstant.SwapOuter).
				TextArea(
					constant.Body,
					"Body",
					editableValue(e.Kind, e.Metadata),
				).
				Lines("3").
				Focus().
				Beside(
					load.New(EventPath(e.Identifier, false), "Cancel").
						Swap(target(e.Identifier), webConstant.SwapOuter).
						Button().
						Class("outline secondary").
						Node(),
				).
				Node(),
		),
	)
}
