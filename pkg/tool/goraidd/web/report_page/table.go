package report_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service/report"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(reports []*report.Report) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(reports))

	for _, r := range reports {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(html.Href(Path(r.Name)), gomponents.Text(r.Name)),
				),
				html.Td(
					html.Class(constant.ReportSizeClass),
					gomponents.Text(size(r.Size)),
				),
				html.Td(
					action.New(
						join.Empty(Path(r.Name), constant.DeleteSuffix),
						"Delete",
					).
						ButtonClass("secondary outline").
						Confirm(fmt.Sprintf("Delete report %s?", r.Name)).
						Small().
						Node(),
				),
			),
		)
	}

	return table.New("Report", "Size", "").
		Rows(rows).
		Empty("No reports generated yet.").
		Node()
}
