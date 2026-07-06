package parse_elite

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"strings"
)

func printBoonTable(players []*elite_parser.AggregatedPlayer) {
	fmt.Println("=== Boon Uptimes (avg %) ===")
	fmt.Printf(
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
	fmt.Println(strings.Repeat("-", 115))

	for _, p := range players {
		f := float64(p.Fights)
		fmt.Printf(
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
