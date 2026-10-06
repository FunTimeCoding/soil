package result

type Constructor struct {
	Type       string   `json:"type"`
	Name       string   `json:"name"`
	File       string   `json:"file"`
	Parameters []string `json:"parameters"`
	Total      int      `json:"total"`
	Inside     int      `json:"inside"`
	Expected   int      `json:"expected"`
	Rewritten  int      `json:"rewritten"`
	Remaining  []*Group `json:"remaining"`
}
