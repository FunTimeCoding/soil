package token_summary

func New(
	statistic []*Statistic,
	block *Spread,
	description *Spread,
	withheld int,
) *Summary {
	return &Summary{
		Statistic:   statistic,
		Block:       block,
		Description: description,
		Withheld:    withheld,
	}
}
