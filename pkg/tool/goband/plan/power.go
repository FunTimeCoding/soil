package plan

type Power struct {
	Instance    string  `json:"instance"`
	Power       string  `json:"power"`
	Refusal     *string `json:"refusal,omitempty"`
	Description string  `json:"description,omitempty"`
}
