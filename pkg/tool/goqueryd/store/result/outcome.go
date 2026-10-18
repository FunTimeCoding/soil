package result

type Outcome struct {
	Results []Search `json:"results"`
	Facets  []Facet  `json:"facets,omitempty"`
}
