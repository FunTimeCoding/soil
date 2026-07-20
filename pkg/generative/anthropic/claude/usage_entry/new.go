package usage_entry

func New(
	timestamp string,
	model string,
	inputTokens int,
	outputTokens int,
	cacheCreationInputTokens int,
	cacheReadInputTokens int,
	cacheCreation5MinuteTokens int,
	cacheCreation1HourTokens int,
) *Entry {
	return &Entry{
		Timestamp:                  timestamp,
		Model:                      model,
		InputTokens:                inputTokens,
		OutputTokens:               outputTokens,
		CacheCreationInputTokens:   cacheCreationInputTokens,
		CacheReadInputTokens:       cacheReadInputTokens,
		CacheCreation5MinuteTokens: cacheCreation5MinuteTokens,
		CacheCreation1HourTokens:   cacheCreation1HourTokens,
	}
}
