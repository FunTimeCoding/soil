package token_summary

func NewStatistic(
	identifier int64,
	name string,
	tags []string,
	block int,
	description int,
	hidden bool,
) *Statistic {
	return &Statistic{
		Identifier:  identifier,
		Name:        name,
		Tags:        tags,
		Block:       block,
		Description: description,
		Hidden:      hidden,
	}
}
