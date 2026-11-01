package counter_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/tally"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(tallies []*tally.Tally) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(tallies))

	for _, v := range tallies {
		cells := []gomponents.Node{html.Td(gomponents.Text(v.Session))}

		for _, n := range []int{
			v.Emitted,
			v.Answered,
			v.Bypassed,
			v.Defaulted,
			v.Declined,
			v.Irrelevant,
			v.Postponed,
			v.Abandoned,
			v.Bumped,
			v.Turns,
		} {
			cells = append(
				cells,
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", n),
				),
			)
		}

		rows = append(rows, html.Tr(cells...))
	}

	return table.New(
		"Session",
		"Emitted",
		"Answered",
		"In chat",
		"Defaulted",
		"Declined",
		"Irrelevant",
		"Postponed",
		"Abandoned",
		"Bumped",
		"Turns",
	).
		Numeric(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
		Rows(rows).
		Empty("Nothing counted yet.").
		Node()
}
