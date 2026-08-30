package example

type Level int

type Payload struct {
	Count   int     `json:"count,omitempty"`
	Ratio   float64 `json:"ratio,omitempty"`
	Enabled bool    `json:"enabled,omitempty"`
	Depth   Level   `json:"depth,omitempty"`
	Name    string  `json:"name,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	Fixed   int     `json:"fixed,omitzero"`
	Items   []int   `json:"items,omitempty"`
	Plain   int     `json:"plain"`
	Skipped int     `json:"-"`
	Bare    int
}
