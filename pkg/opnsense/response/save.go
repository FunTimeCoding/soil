package response

type Save struct {
	Result     string            `json:"result"`
	Identifier string            `json:"uuid"`
	Validation map[string]string `json:"validations"`
}
