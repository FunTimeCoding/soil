package parse_elite

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"strings"
)

func printSupportTable(players []*elite_parser.AggregatedPlayer) {
	console.Line("=== Support ===")
	console.Format(
		"%-25s %-15s %6s %6s %6s %8s %8s %5s %5s\n",
		"Name",
		"Profession",
		"Fights",
		"Strips",
		"Cleans",
		"Heal",
		"Barrier",
		"Res",
		"Stun",
	)
	console.Line(strings.Repeat("-", 110))

	for _, p := range players {
		console.Format(
			"%-25s %-15s %6d %6d %6d %8d %8d %5d %5d\n",
			p.Name,
			p.Profession,
			p.Fights,
			p.Support.BoonStrips,
			p.Support.ConditionCleanses,
			p.Support.Healing,
			p.Support.Barrier,
			p.Support.Resurrects,
			p.Support.StunBreak,
		)
	}
}
