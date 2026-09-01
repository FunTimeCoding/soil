package parse_elite

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"strings"
)

func printDefensiveTable(players []*elite_parser.AggregatedPlayer) {
	console.Line("=== Defensive ===")
	console.Format(
		"%-25s %-15s %6s %6s %6s %6s %6s %6s %6s\n",
		"Name",
		"Profession",
		"Fights",
		"Deaths",
		"Downed",
		"Dodge",
		"Evade",
		"Block",
		"Invuln",
	)
	console.Line(strings.Repeat("-", 100))

	for _, p := range players {
		console.Format(
			"%-25s %-15s %6d %6d %6d %6d %6d %6d %6d\n",
			p.Name,
			p.Profession,
			p.Fights,
			p.Defensive.DeadCount,
			p.Defensive.DownCount,
			p.Defensive.DodgeCount,
			p.Defensive.EvadedCount,
			p.Defensive.BlockedCount,
			p.Defensive.InvulnedCount,
		)
	}
}
