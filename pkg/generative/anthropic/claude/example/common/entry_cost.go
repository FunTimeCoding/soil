package common

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/usage_entry"

func EntryCost(
	model string,
	e *usage_entry.Entry,
) float64 {
	var inputRate float64
	var outputRate float64
	var cacheCreateRate float64
	var cacheReadRate float64

	switch model {
	case "opus":
		inputRate = 5.0
		outputRate = 25.0
		cacheCreateRate = 6.25
		cacheReadRate = 0.5
	case "sonnet":
		inputRate = 3.0
		outputRate = 15.0
		cacheCreateRate = 3.75
		cacheReadRate = 0.3
	case "haiku":
		inputRate = 1.0
		outputRate = 5.0
		cacheCreateRate = 1.25
		cacheReadRate = 0.1
	default:
		inputRate = 3.0
		outputRate = 15.0
		cacheCreateRate = 3.75
		cacheReadRate = 0.3
	}

	return float64(e.InputTokens)/1_000_000*inputRate +
		float64(e.OutputTokens)/1_000_000*outputRate +
		float64(e.CacheCreationInputTokens)/1_000_000*
			cacheCreateRate +
		float64(e.CacheReadInputTokens)/1_000_000*
			cacheReadRate
}
