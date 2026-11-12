package response

type Limit struct {
	Kind     string  `json:"kind"`
	Percent  int     `json:"percent"`
	ResetsAt *string `json:"resets_at"`
	Scope    *Scope  `json:"scope"`
}
