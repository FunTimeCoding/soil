package search

type ListOutcome struct {
	Results []Result `json:"results"`
	Facets  []Facet  `json:"facets,omitempty"`
}
