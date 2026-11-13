package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func opening(e *entry.Entry) *opening_row.OpeningRow {
	return opening_row.New(
		rowMark(e.Identifier),
		fragment(constant.DetailPath, e.Identifier),
		fragment(constant.ClosePath, e.Identifier),
	).Cells(
		layout.TimeCell(e.Timestamp),
		html.Td(gomponents.Text(e.Action)),
		html.Td(gomponents.Text(e.User)),
		html.Td(gomponents.Text(e.System)),
		html.Td(gomponents.Text(e.Service)),
		layout.ClampCell(e.Description),
	)
}
