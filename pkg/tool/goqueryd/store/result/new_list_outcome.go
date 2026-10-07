package result

func NewListOutcome(
	results []Search,
	facets []Facet,
) *ListOutcome {
	return &ListOutcome{Results: results, Facets: facets}
}
