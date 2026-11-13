package opening_row

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func (r *OpeningRow) row(extra ...gomponents.Node) gomponents.Node {
	path := r.open

	if r.expanded {
		path = r.close
	}

	return html.Tr(
		html.ID(r.identifier),
		html.Class(constant.OpeningRowClass),
		gomponents.Attr(
			constant.ExpandedAttribute,
			strconv.FormatBool(r.expanded),
		),
		extended.Get(path),
		extended.Target(r.Target()),
		extended.Swap(constant.SwapOuter),
		gomponents.Group(extra),
		gomponents.Group(r.cells),
	)
}
