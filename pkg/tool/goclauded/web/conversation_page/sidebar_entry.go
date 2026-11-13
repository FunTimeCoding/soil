package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/enriched_session"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sidebarEntry(
	s *enriched_session.Session,
	current bool,
) gomponents.Node {
	name := s.Slug

	if s.Alias != "" {
		name = s.Alias
	}

	if name == "" {
		name = constant.UnnamedSession
	}

	return html.Div(
		html.ID(join.Empty(constant.EntryPrefix, s.Identifier)),
		html.Class(constant.SidebarEntryClass),
		gomponents.If(
			current,
			gomponents.Attr(webConstant.CurrentAttribute, "true"),
		),
		html.Div(
			html.Class(constant.EntryNameClass),
			html.A(
				html.Class(constant.EntryLinkClass),
				html.Href(panelPath(s.Identifier)),
				extended.Get(panelPath(s.Identifier)),
				extended.Target(constant.PanelTarget),
				extended.Push("true"),
				gomponents.Text(name),
			),
			load.New(editPath(s.Identifier), "✎").
				Swap(constant.PanelTarget, webConstant.SwapInner).
				Class(constant.RenameIconClass).
				Node(),
		),
		gomponents.If(
			s.Description != "",
			html.Small(
				html.Class(webConstant.CardClampClass),
				gomponents.Text(s.Description),
			),
		),
		html.Small(gomponents.Text(relativeTimestamp(s.Timestamp))),
	)
}
