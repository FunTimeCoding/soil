package usage

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/common"
	"sort"
	"time"
)

func Debug() {
	c := claude.New()
	sessions := c.Sessions()
	var all []*common.Timestamped

	for _, s := range sessions {
		for _, entry := range c.UsageEntries(s.Identifier) {
			t := common.ParseTimestamp(entry.Timestamp)

			if t.IsZero() {
				continue
			}

			all = append(all, common.New(t, entry))
		}
	}

	sort.Slice(
		all,
		func(i, j int) bool {
			return all[i].Time.Before(all[j].Time)
		},
	)
	fmt.Printf("Total entries across all sessions: %d\n\n", len(all))
	window := 5 * time.Hour
	fmt.Println("Gaps >= 1h in last 24h:")
	cutoff24 := now().Add(-24 * time.Hour)

	for i := 1; i < len(all); i++ {
		if all[i].Time.Before(cutoff24) {
			continue
		}

		gap := all[i].Time.Sub(all[i-1].Time)

		if gap >= time.Hour {
			marker := ""

			if gap >= window {
				marker = " *** BLOCK BREAK ***"
			}

			fmt.Printf(
				"  %s → %s  gap=%s%s\n",
				all[i-1].Time.Format("Jan 02 15:04"),
				all[i].Time.Format("15:04"),
				gap.Truncate(time.Second),
				marker,
			)
		}
	}

	fmt.Println()
	block := findActiveBlock(all)

	if block == nil {
		fmt.Println("no active block")

		return
	}

	fmt.Printf(
		"Active block: %s - %s (%d entries)\n",
		block.start.Format("15:04"),
		block.end.Format("15:04"),
		len(block.entries),
	)
	var totalCost float64

	for _, ts := range block.entries {
		totalCost += common.EntryCost(
			common.NormalizeModel(ts.Entry.Model),
			ts.Entry,
		)
	}

	fmt.Printf(
		"Cost: $%.2f / $140.00  (%.0f%%)\n",
		totalCost,
		(totalCost/140)*100,
	)
}
