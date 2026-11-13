package opening_row

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (r *OpeningRow) Removal() gomponents.Node {
	return html.Template(
		html.Tr(html.ID(r.identifier), extended.OutOfBand(constant.SwapDelete)),
	)
}
