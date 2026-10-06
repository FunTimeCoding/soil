package search

func NewListOutcome(
	results []Result,
	facets []Facet,
) *ListOutcome {
	return &ListOutcome{Results: results, Facets: facets}
}
