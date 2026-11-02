package result

type Literals struct {
	Type     string   `json:"type"`
	Total    int      `json:"total"`
	Inside   int      `json:"inside"`
	Expected int      `json:"expected"`
	Groups   []*Group `json:"groups"`
}
