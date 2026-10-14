package fact

type Summary struct {
	Arranges  []string               `json:"arranges"`
	Arranging map[string][]Condition `json:"arranging"`
	Delegates map[string][]string    `json:"delegates"`
	Disposes  map[string][]int       `json:"disposes"`
	Forwards  map[string][]Forward   `json:"forwards"`
	Releases  map[string][]Release   `json:"releases"`
	Shared    []string               `json:"shared"`
	Lends     []string               `json:"lends"`
	Self      []string               `json:"self"`
	Spent     []string               `json:"spent"`
	SpentVia  map[string][]string    `json:"spent_via"`
	LendsVia  map[string][]string    `json:"lends_via"`
	Wraps     map[string][]Forward   `json:"wraps"`
}
