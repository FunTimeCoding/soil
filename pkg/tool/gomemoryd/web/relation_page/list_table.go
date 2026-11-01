package relation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func listTable(
	relations []record.RelationOverview,
	filter string,
) gomponents.Node {
	var rows []gomponents.Node

	for _, v := range relations {
		kind := v.Type
		label := v.Type

		if kind == "" {
			kind = constant.UntypedFilter
			label = "-"
		}

		if filter != "" && kind != filter {
			continue
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(memory_page.Link(v.SourceIdentifier, v.SourceName)),
				scopeCell(v.SourceScope),
				html.Td(gomponents.Text(label)),
				html.Td(memory_page.Link(v.TargetIdentifier, v.TargetName)),
				scopeCell(v.TargetScope),
			),
		)
	}

	return table.New("Source", "Scope", "Type", "Target", "Scope").
		Rows(rows).
		Empty("No relations.").
		Node()
}
