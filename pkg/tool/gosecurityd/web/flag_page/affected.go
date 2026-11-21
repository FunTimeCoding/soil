package flag_page

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func Affected(a *affected_result.Result) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(a.Rows))

	for _, r := range a.Rows {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(
							join.Empty(
								constant.SystemPath,
								"?",
								url.Values{
									generativeConstant.ParameterName: {r.System},
								}.Encode(),
							),
						),
						gomponents.Text(r.System),
					),
				),
				html.Td(system_page.CoverageBadge(r.Coverage)),
				html.Td(gomponents.Text(r.Source)),
				html.Td(html.Small(gomponents.Text(r.Installed))),
				html.Td(html.Small(gomponents.Text(r.Fixed))),
				html.Td(html.Small(gomponents.Text(r.Candidate))),
				html.Td(gomponents.Text(r.Kind)),
			),
		)
	}

	return html.Section(
		html.ID(constant.SystemsMark),
		html.H2(gomponents.Text("Affected systems")),
		html.P(gomponents.Text(hitSummary(a.Count))),
		table.New(
			"System",
			"Coverage",
			"Package",
			"Installed",
			"Fixed in",
			"Candidate",
			"Fix",
		).
			Rows(rows).
			Empty("No reporting system carries it.").
			Node(),
	)
}
