package search

type Outcome struct {
	Results  []Result `json:"results"`
	Facets   []Facet  `json:"facets,omitempty"`
	Degraded bool     `json:"degraded,omitzero"`
	Cause    error    `json:"-"`
}
