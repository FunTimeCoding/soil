package entry_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(entries []entry.Entry) gomponents.Node {
	rows := make([]gomponents.Node, 0, 2*len(entries))

	for _, e := range entries {
		mark := detailMark(e.Identifier)
		rows = append(
			rows,
			html.Tr(
				html.ID(fmt.Sprintf("%s%d", constant.RowPrefix, e.Identifier)),
				html.Class(constant.ClickableRowClass),
				extended.Get(fragment(constant.DetailPath, e.Identifier)),
				extended.Target(join.Empty("#", mark)),
				extended.Swap(webConstant.SwapOuter),
				layout.TimeCell(e.Timestamp),
				html.Td(gomponents.Text(e.Action)),
				html.Td(gomponents.Text(e.User)),
				html.Td(gomponents.Text(e.System)),
				html.Td(gomponents.Text(e.Service)),
				html.Td(gomponents.Text(truncate(e.Description, 80))),
			),
			html.Tr(html.ID(mark), gomponents.Attr("hidden")),
		)
	}

	return table.New(
		"Timestamp",
		"Action",
		"User",
		"System",
		"Service",
		"Description",
	).
		Rows(rows).
		Empty("No entries found.").
		Node()
}
