package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Layout(
	query string,
	kinds []string,
	sidebar gomponents.Node,
	panel gomponents.Node,
) gomponents.Node {
	if panel == nil {
		panel = html.P(
			html.Class(constant.PanelPlaceholderClass),
			gomponents.Text("Select a conversation"),
		)
	}

	return html.Div(
		html.Class(constant.ConversationLayoutClass),
		html.Div(
			html.Class(constant.SidebarClass),
			searchForm(query, kinds),
			html.Div(
				html.ID(constant.SidebarEntriesIdentifier),
				extended.Get(sidebarPath(0)),
				extended.Trigger(
					join.Empty(constant.SessionEditedEvent, " from:body"),
				),
				extended.Swap(webConstant.SwapInner),
				sidebar,
			),
		),
		html.Div(
			html.Class(constant.PanelClass),
			html.ID(constant.PanelIdentifier),
			panel,
		),
		html.Script(gomponents.Raw(constant.PanelSwapScript)),
		html.Script(gomponents.Raw(constant.HitNavigationScript)),
	)
}
