package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Missing() gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H1(gomponents.Text(constant.EntryTitle)),
			html.P(
				html.Class(webConstant.AbsentClass),
				gomponents.Text("This entry no longer exists."),
			),
			back(),
		},
	)
}
