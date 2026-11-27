package request

type KeyActions struct {
	Type       string `json:"type"`
	Identifier string `json:"id"`
	Actions    []*Key `json:"actions"`
}
