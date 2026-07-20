package usage

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/common"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"sort"
)

func Usage() {
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
	activeBlock := findActiveBlock(all)

	if activeBlock == nil {
		fmt.Println("no active 5h block")

		return
	}

	type modelTotals struct {
		inputTokens              int
		outputTokens             int
		cacheCreationInputTokens int
		cacheReadInputTokens     int
		count                    int
	}
	byModel := map[string]*modelTotals{}
	var totalCost float64

	for _, t := range activeBlock.entries {
		key := pricing.NormalizeModel(t.Entry.Model)
		m, okay := byModel[key]

		if !okay {
			m = &modelTotals{}
			byModel[key] = m
		}

		m.inputTokens += t.Entry.InputTokens
		m.outputTokens += t.Entry.OutputTokens
		m.cacheCreationInputTokens += t.Entry.CacheCreationInputTokens
		m.cacheReadInputTokens += t.Entry.CacheReadInputTokens
		m.count++
		totalCost += pricing.EntryCost(key, t.Entry)
	}

	costLimit := 140.0
	percentage := (totalCost / costLimit) * 100
	remaining := activeBlock.end.Sub(now())

	if remaining < 0 {
		remaining = 0
	}

	fmt.Printf(
		"5h block: %s - %s  (%s remaining)\n\n",
		activeBlock.start.Format("15:04"),
		activeBlock.end.Format("15:04"),
		formatDuration(remaining),
	)

	for model, m := range byModel {
		fmt.Printf(
			"  %-12s  %6d input, %6d output, %6d cache-create, %6d cache-read  (%d calls)\n",
			model,
			m.inputTokens,
			m.outputTokens,
			m.cacheCreationInputTokens,
			m.cacheReadInputTokens,
			m.count,
		)
	}

	fmt.Printf(
		"\nCost: $%.2f / $%.2f  (%.0f%%)\n",
		totalCost,
		costLimit,
		percentage,
	)
}
