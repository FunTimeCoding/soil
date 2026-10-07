package result

type Facet struct {
	Key      string         `json:"key"`
	Distinct int            `json:"distinct"`
	Values   map[string]int `json:"values,omitempty"`
}
