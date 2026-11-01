package intent

type Change struct {
	Name string `json:"name"`
	Node string `json:"node"`
	From string `json:"from"`
	To   string `json:"to"`
}
