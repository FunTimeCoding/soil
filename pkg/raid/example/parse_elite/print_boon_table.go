package parse_elite

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"strings"
)

func printBoonTable(players []*elite_parser.AggregatedPlayer) {
	console.Line("=== Boon Uptimes (avg %) ===")
	console.Format(
		"%-25s %-15s %6s %6s %6s %6s %6s %6s %6s %6s\n",
		"Name",
		"Profession",
		"Fights",
		"Stab",
		"Might",
		"Fury",
		"Prot",
		"Quick",
		"Resist",
		"Alac",
	)
	console.Line(strings.Repeat("-", 115))

	for _, p := range players {
		f := float64(p.Fights)
		console.Format(
			"%-25s %-15s %6d %5.1f%% %5.1f%% %5.1f%% %5.1f%% %5.1f%% %5.1f%% %5.1f%%\n",
			p.Name,
			p.Profession,
			p.Fights,
			p.Boons.Stability/f,
			p.Boons.Might/f,
			p.Boons.Fury/f,
			p.Boons.Protection/f,
			p.Boons.Quickness/f,
			p.Boons.Resistance/f,
			p.Boons.Alacrity/f,
		)
	}
}
