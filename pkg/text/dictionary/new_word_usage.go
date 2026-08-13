package dictionary

func NewWordUsage(
	word string,
	category string,
	used bool,
) *WordUsage {
	return &WordUsage{Word: word, Category: category, Used: used}
}
