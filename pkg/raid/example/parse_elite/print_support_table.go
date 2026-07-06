package parse_elite

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"strings"
)

func printSupportTable(players []*elite_parser.AggregatedPlayer) {
	fmt.Println("=== Support ===")
	fmt.Printf(
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
	fmt.Println(strings.Repeat("-", 110))

	for _, p := range players {
		fmt.Printf(
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
