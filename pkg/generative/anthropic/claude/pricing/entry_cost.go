package pricing

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/usage_entry"

func EntryCost(
	model string,
	e *usage_entry.Entry,
) float64 {
	return Cost(
		model,
		&Tokens{
			Input:         e.InputTokens,
			Output:        e.OutputTokens,
			CacheCreation: e.CacheCreationInputTokens,
			CacheRead:     e.CacheReadInputTokens,
			Cache5Minute:  e.CacheCreation5MinuteTokens,
			Cache1Hour:    e.CacheCreation1HourTokens,
		},
	)
}
