package pricing

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/usage_entry"

func (t *Tokens) AddEntry(e *usage_entry.Entry) {
	t.Input += e.InputTokens
	t.Output += e.OutputTokens
	t.CacheCreation += e.CacheCreationInputTokens
	t.CacheRead += e.CacheReadInputTokens
	t.Cache5Minute += e.CacheCreation5MinuteTokens
	t.Cache1Hour += e.CacheCreation1HourTokens
	t.Count++
}
