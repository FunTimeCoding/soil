package player_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func raidTable(rows []record.PlayerRaid) gomponents.Node {
	nodes := make([]gomponents.Node, 0, len(rows))

	for _, r := range rows {
		seconds := float64(r.ActiveTimeMS) / 1000
		minutes := seconds / 60
		nodes = append(
			nodes,
			html.Tr(
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
		Numeric(1, 2, 3, 4, 5, 6, 7, 8, 9).
		Rows(nodes).
		Node()
}
