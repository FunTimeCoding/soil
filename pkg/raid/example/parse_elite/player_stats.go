package parse_elite

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/raid/elite"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"os"
	"path/filepath"
	"strings"
)

func PlayerStats() {
	a := argument.NewSimple("player-stats")
	a.ParseSimple()
	directory := a.RequiredPositional(0, "DIRECTORY")
	entries, e := os.ReadDir(directory)
	errors.PanicOnError(e)
	var allStats []*elite_parser.PlayerFightStatistic
	fightCount := 0
	skipped := 0

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), constant.DetailedWvWKillSuffix) {
			continue
		}

		path := filepath.Join(directory, entry.Name())
		b, readError := os.ReadFile(path)
		errors.PanicOnError(readError)
		var fight elite.Fight
		errors.PanicOnError(json.Unmarshal(b, &fight))

		if !elite_parser.IsValidFight(&fight) {
			skipped++

			continue
		}

		fightCount++
		allStats = append(allStats, elite_parser.Extract(&fight)...)
	}

	players := elite_parser.Aggregate(allStats)
	fmt.Printf(
		"Aggregated %d fights (%d skipped), %d player-profession entries\n\n",
		fightCount,
		skipped,
		len(players),
	)
	sorted := sortByDamage(players)
	printOffensiveTable(sorted)
	fmt.Println()
	printDefensiveTable(sorted)
	fmt.Println()
	printSupportTable(sorted)
	fmt.Println()
	printBoonTable(sorted)
}
