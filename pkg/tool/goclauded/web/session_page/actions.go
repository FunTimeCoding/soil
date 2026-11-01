package session_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func (p *Page) actions() gomponents.Node {
	identifier := p.view.Detail.Identifier

	return gomponents.Group(
		[]gomponents.Node{
			html.P(
				gomponents.Text("Delete confirmation: "),
				html.Code(gomponents.Text(p.view.DeleteHash)),
			),
			html.Div(
				html.Class(webConstant.ClusterClass),
				buttonLink(ActionPath(identifier, constant.EditAction), "Edit"),
				buttonLink(
					join.Empty(
						constant.ConversationsPath,
						stringsConstant.Slash,
						url.PathEscape(identifier),
					),
					"View conversation",
				),
				action.New(
					ActionPath(identifier, constant.DeleteAction),
					"Delete",
				).
					Confirm("Delete this session and all its data?").
					ButtonClass("outline contrast").
					Node(),
			),
		},
	)
}
