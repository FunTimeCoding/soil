package opening_row

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func (r *OpeningRow) detailRow() gomponents.Node {
	if !r.expanded {
		return html.Tr(html.ID(r.detailIdentifier()), html.Hidden("hidden"))
	}

	return html.Tr(
		html.ID(r.detailIdentifier()),
		html.Class(constant.OpeningDetailClass),
		html.Td(
			html.ColSpan(strconv.Itoa(r.columns)),
			html.Div(
				html.Class(constant.OpeningContentClass),
				gomponents.Group(r.detail),
			),
		),
	)
}
