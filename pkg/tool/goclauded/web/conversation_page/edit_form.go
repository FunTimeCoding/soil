package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func EditForm(
	identifier string,
	alias string,
	description string,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H4(gomponents.Text("Edit Session")),
			entry.New(editPath(identifier), "Save").
				Swap(constant.PanelTarget, webConstant.SwapInner).
				Sideways().
				Text(constant.Alias, "Name", alias).
				Focus().
				TextArea(constant.Description, "Description", description).
				Lines("3").
				Beside(
					load.New(panelPath(identifier), "Cancel").
						Swap(constant.PanelTarget, webConstant.SwapInner).
						Button().
						Class("outline secondary").
						Node(),
				).
				Node(),
		},
	)
}
