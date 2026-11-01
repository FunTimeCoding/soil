package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Permalink(e *entry.Entry) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H1(
				gomponents.Textf("%s %d", constant.EntryTitle, e.Identifier),
			),
			html.Div(html.Class(constant.DetailContentClass), fields(e)),
			back(),
		},
	)
}
