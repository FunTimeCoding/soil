package result

type ListOutcome struct {
	Results []Search `json:"results"`
	Facets  []Facet  `json:"facets,omitempty"`
}
