package store

type SearchOutcome struct {
	Results  []SearchResult `json:"results"`
	Facets   []Facet        `json:"facets,omitempty"`
	Degraded bool           `json:"degraded,omitzero"`
	Cause    error          `json:"-"`
}
