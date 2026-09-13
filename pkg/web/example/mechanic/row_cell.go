package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func rowCell(replaced bool) gomponents.Node {
	text := "original"

	if replaced {
		text = "replaced"
	}

	return html.Div(html.ID(constant.RowMark), gomponents.Text(text))
}
