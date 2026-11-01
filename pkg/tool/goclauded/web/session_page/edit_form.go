package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func EditForm(s *session.Session) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Edit Session")),
			entry.New(ActionPath(s.Identifier, constant.EditAction), "Save").
				Record().
				Text(constant.Alias, "Name", s.AliasValue()).
				TextArea(constant.Description, "Description", s.Description).
				Lines("3").
				Beside(
					html.A(
						html.Href(Path(s.Identifier)),
						gomponents.Text("Cancel"),
					),
				).
				Node(),
		},
	)
}
