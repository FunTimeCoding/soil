package usage

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/common"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/usage_entry"
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
		console.Line("no active 5h block")

		return
	}

	byModel := map[string]*usage_entry.Entry{}
	calls := map[string]int{}
	var totalCost float64

	for _, t := range activeBlock.entries {
		key := pricing.NormalizeModel(t.Entry.Model)
		m, okay := byModel[key]

		if !okay {
			m = usage_entry.New("", key, 0, 0, 0, 0, 0, 0)
			byModel[key] = m
		}

		m.InputTokens += t.Entry.InputTokens
		m.OutputTokens += t.Entry.OutputTokens
		m.CacheCreationInputTokens += t.Entry.CacheCreationInputTokens
		m.CacheReadInputTokens += t.Entry.CacheReadInputTokens
		calls[key]++
		totalCost += pricing.EntryCost(key, t.Entry)
	}

	costLimit := 140.0
	percentage := (totalCost / costLimit) * 100
	remaining := activeBlock.end.Sub(now())

	if remaining < 0 {
		remaining = 0
	}

	console.Format(
		"5h block: %s - %s  (%s remaining)\n\n",
		activeBlock.start.Format("15:04"),
		activeBlock.end.Format("15:04"),
		formatDuration(remaining),
	)

	for model, m := range byModel {
		console.Format(
			"  %-12s  %6d input, %6d output, %6d cache-create, %6d cache-read  (%d calls)\n",
			model,
			m.InputTokens,
			m.OutputTokens,
			m.CacheCreationInputTokens,
			m.CacheReadInputTokens,
			calls[model],
		)
	}

	console.Format(
		"\nCost: $%.2f / $%.2f  (%.0f%%)\n",
		totalCost,
		costLimit,
		percentage,
	)
}
