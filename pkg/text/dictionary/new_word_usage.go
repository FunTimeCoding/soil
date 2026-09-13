package dictionary

import "strings"

func NewWordUsage(
	word string,
	category string,
	used bool,
) *WordUsage {
	return &WordUsage{
		Word:     word,
		Category: category,
		Used:     used,
		lower:    strings.ToLower(word),
	}
}
