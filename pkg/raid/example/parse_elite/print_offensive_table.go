package parse_elite

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"strings"
)

func printOffensiveTable(players []*elite_parser.AggregatedPlayer) {
	fmt.Println("=== Offensive ===")
	fmt.Printf(
		"%-25s %-15s %6s %10s %6s %6s %6s\n",
		"Name",
		"Profession",
		"Fights",
		"Damage",
		"DPS",
		"Kills",
		"Downs",
	)
	fmt.Println(strings.Repeat("-", 90))

	for _, p := range players {
		dps := 0

		if p.TotalActiveTimeMS > 0 {
			dps = p.Offensive.Damage * 1000 / p.TotalActiveTimeMS
		}

		fmt.Printf(
			"%-25s %-15s %6d %10d %6d %6d %6d\n",
			p.Name,
			p.Profession,
			p.Fights,
			p.Offensive.Damage,
			dps,
			p.Offensive.Kills,
			p.Offensive.Downs,
		)
	}
}
