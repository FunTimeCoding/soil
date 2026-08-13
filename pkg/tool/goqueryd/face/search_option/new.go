package search_option

func New(
	query string,
	collection string,
	limit int,
) *Option {
	return &Option{Query: query, Collection: collection, Limit: limit}
}
