package utilization

type limit struct {
	Kind     string  `json:"kind"`
	Percent  int     `json:"percent"`
	ResetsAt *string `json:"resets_at"`
	Scope    *scope  `json:"scope"`
}
