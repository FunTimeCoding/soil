package update

type Result struct {
	Changed bool   `json:"changed"`
	Diff    string `json:"diff,omitempty"`
	Error   error  `json:"-"`
}
