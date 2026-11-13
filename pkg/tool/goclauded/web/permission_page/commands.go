package permission_page

import (
	"github.com/funtimecoding/soil/pkg/strings/shorten"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func commands(w *permission.Row) gomponents.Node {
	result := make([]gomponents.Node, 0, len(w.Commands))

	for _, c := range w.Commands {
		result = append(
			result,
			html.Pre(
				gomponents.Text(
					shorten.Ellipsis(c, constant.PermissionCommandLength),
				),
			),
		)
	}

	return gomponents.Group(result)
}
