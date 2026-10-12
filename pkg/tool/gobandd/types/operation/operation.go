package operation

type Operation struct {
	Name    string `json:"name"`
	Force   bool   `json:"force"`
	State   string `json:"state"`
	Refusal string `json:"refusal,omitempty"`
}
