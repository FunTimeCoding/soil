package request

type Actions struct {
	Context string       `json:"context"`
	Actions []*KeyActions `json:"actions"`
}
