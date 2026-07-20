package usage_entry

type Entry struct {
	Timestamp                  string
	Model                      string
	InputTokens                int
	OutputTokens               int
	CacheCreationInputTokens   int
	CacheReadInputTokens       int
	CacheCreation5MinuteTokens int
	CacheCreation1HourTokens   int
}
