package result

func NewApply(
	symbol string,
	pattern string,
	replacement string,
) *Apply {
	return &Apply{
		Symbol:      symbol,
		Pattern:     pattern,
		Replacement: replacement,
	}
}
