package choice_set_request

type Request struct {
	Name    string     `json:"name,omitempty"`
	Choices [][]string `json:"extra_choices"`
}
