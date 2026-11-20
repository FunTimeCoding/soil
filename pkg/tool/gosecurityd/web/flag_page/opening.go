package flag_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func opening(r *overview_result.Row) *opening_row.OpeningRow {
	name := r.Flag.Vulnerability

	return opening_row.New(
		join.Empty(constant.RowPrefix, name),
		fragment(constant.FlagOpenPath, name),
		fragment(constant.FlagShutPath, name),
	).Cells(
		html.Td(gomponents.Text(name)),
		html.Td(SeverityBadge(r.Flag)),
		html.Td(DebianChips(r)),
		layout.ClampCell(r.Flag.Note),
		layout.TimeCell(r.Flag.CreatedAt),
	)
}
