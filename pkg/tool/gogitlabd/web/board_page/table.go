package board_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/component"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/pipeline_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func Table(
	entries []*board_entry.Entry,
	multiple bool,
) gomponents.Node {
	prefix := namespace(entries)
	var columns []string

	if multiple {
		columns = append(columns, "Instance")
	}

	columns = append(columns, "", "Project", "Branch", "Pipeline", "Updated")
	rows := make([]gomponents.Node, 0, len(entries))

	for _, e := range entries {
		rows = append(
			rows,
			html.Tr(
				gomponents.If(
					multiple,
					html.Td(
						html.Class(constant.ReferenceClass),
						gomponents.Text(e.Instance),
					),
				),
				html.Td(component.StatusIcon(e.Status)),
				html.Td(
					html.A(
						html.Href(e.ProjectLink),
						html.Target("_blank"),
						gomponents.Text(strings.TrimPrefix(e.Project, prefix)),
					),
				),
				html.Td(
					html.Class(constant.ReferenceClass),
					gomponents.Text(
						strings.TrimPrefix(e.Reference, constant.RenovatePrefix),
					),
				),
				html.Td(
					html.A(
						html.Href(
							pipeline_page.Path(
								e.Instance,
								e.ProjectIdentifier,
								e.Identifier,
								0,
							),
						),
						gomponents.Text(fmt.Sprintf("#%d", e.Identifier)),
					),
				),
				layout.TimeCell(e.Updated),
			),
		)
	}

	return table.New(columns...).
		Class(constant.BoardClass).
		Rows(rows).
		Empty("No pipelines seen yet - first poll pending.").
		Node()
}
