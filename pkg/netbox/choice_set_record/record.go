package choice_set_record

type Record struct {
	Identifier  int32      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Choices     [][]string `json:"extra_choices"`
}
