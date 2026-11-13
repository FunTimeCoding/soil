package opening_row

import (
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (r *OpeningRow) Response() gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			r.detailRow(),
			html.Template(r.row(extended.OutOfBand("true"))),
		},
	)
}
