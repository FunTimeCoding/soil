package result

func NewMatch(
	symbol string,
	pattern string,
	total int,
	matched int,
	unmatched []*Group,
) *Match {
	return &Match{
		Symbol:    symbol,
		Pattern:   pattern,
		Total:     total,
		Matched:   matched,
		Unmatched: unmatched,
	}
}
