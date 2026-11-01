package raid_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(players []record.RaidPlayer) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(players))

	for _, r := range players {
		seconds := float64(r.ActiveTimeMS) / 1000
		minutes := seconds / 60
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(component.PlayerPath(r.Account)),
						gomponents.Text(r.Name),
					),
				),
				component.ProfessionCell(r.Profession),
				component.NumericCell(fmt.Sprintf("%d", r.Fights)),
				component.NumericCell(component.PerSecond(r.Damage, seconds, 1)),
				component.NumericCell(
					component.PerSecond(r.Healing, seconds, 1),
				),
				component.NumericCell(
					component.PerSecond(r.ConditionCleanses, seconds, 1),
				),
				component.NumericCell(
					component.PerSecond(r.BoonStrips, seconds, 2),
				),
				component.NumericCell(
					component.PerSecond(r.Barrier, seconds, 1),
				),
				component.NumericCell(component.PerSecond(r.Downs, seconds, 1)),
				component.NumericCell(component.PerMinute(r.DeadCount, minutes)),
				component.NumericCell(fmt.Sprintf("%.0f", r.DistToCom)),
			),
		)
	}

	return table.New(
		"Player",
		"Profession",
		"Fights",
		"Dmg/s",
		"Heal/s",
		"Clean/s",
		"Strip/s",
		"Barrier/s",
		"Downs/s",
		"Death/m",
		"Tag Dist",
	).
		Numeric(2, 3, 4, 5, 6, 7, 8, 9, 10).
		Rows(rows).
		Empty("No player stats for this raid.").
		Node()
}
