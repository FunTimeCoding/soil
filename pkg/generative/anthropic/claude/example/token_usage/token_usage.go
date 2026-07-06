package token_usage

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/common"
	library "github.com/funtimecoding/soil/pkg/time"
	"sort"
	"time"
)

func TokenUsage(
	from time.Time,
	to time.Time,
) {
	c := claude.New()
	sessions := c.Sessions()
	var all []*common.Timestamped

	for _, s := range sessions {
		for _, entry := range c.UsageEntries(s.Identifier) {
			t := common.ParseTimestamp(entry.Timestamp)

			if t.IsZero() {
				continue
			}

			if t.Before(from) || !t.Before(to) {
				continue
			}

			all = append(all, common.New(t, entry))
		}
	}

	if len(all) == 0 {
		fmt.Printf(
			"No token usage between %s and %s\n",
			from.Format(time.RFC3339),
			to.Format(time.RFC3339),
		)

		return
	}

	sort.Slice(
		all,
		func(i, j int) bool {
			return all[i].Time.Before(all[j].Time)
		},
	)
	byModel := map[string]*modelTotals{}
	var totalCost float64

	for _, ts := range all {
		key := common.NormalizeModel(ts.Entry.Model)
		m, okay := byModel[key]

		if !okay {
			m = &modelTotals{}
			byModel[key] = m
		}

		m.input += ts.Entry.InputTokens
		m.output += ts.Entry.OutputTokens
		m.cacheCreate += ts.Entry.CacheCreationInputTokens
		m.cacheRead += ts.Entry.CacheReadInputTokens
		m.count++
		totalCost += common.EntryCost(key, ts.Entry)
	}

	fmt.Printf(
		"Token usage: %s → %s\n",
		from.Format(library.DateMinute),
		to.Format(library.DateMinute),
	)
	fmt.Printf("Entries: %d\n\n", len(all))

	for model, m := range byModel {
		if m.input+m.output+m.cacheCreate+m.cacheRead == 0 {
			continue
		}

		fmt.Printf("  %-12s  %d calls\n", model, m.count)
		fmt.Printf(
			"    input:        %10d\n",
			m.input,
		)
		fmt.Printf(
			"    output:       %10d\n",
			m.output,
		)
		fmt.Printf(
			"    cache-create: %10d\n",
			m.cacheCreate,
		)
		fmt.Printf(
			"    cache-read:   %10d\n",
			m.cacheRead,
		)
		fmt.Printf(
			"    cost:         $%.2f\n\n",
			modelCost(model, m),
		)
	}

	fmt.Printf("Total cost: $%.2f\n", totalCost)
}
