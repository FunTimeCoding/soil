package statistic_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Region(summary *token_summary.Summary) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.StatisticTitle)),
			html.P(gomponents.Text(constant.StatisticIntroduction)),
			spread(summary),
			listing(summary),
			gomponents.If(
				summary.Withheld > 0,
				html.P(
					gomponents.Textf(
						"%d withheld from this listing; the spreads include them.",
						summary.Withheld,
					),
				),
			),
		},
	)
}
