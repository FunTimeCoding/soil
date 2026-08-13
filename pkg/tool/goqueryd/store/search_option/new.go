package search_option

func New(
	query string,
	limit int,
) *Option {
	return &Option{Query: query, Limit: limit}
}
