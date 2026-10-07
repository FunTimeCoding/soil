package result

type Outcome struct {
	Results  []Search `json:"results"`
	Facets   []Facet  `json:"facets,omitempty"`
	Degraded bool     `json:"degraded,omitzero"`
	Cause    error    `json:"-"`
}
