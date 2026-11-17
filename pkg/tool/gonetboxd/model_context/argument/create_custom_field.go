package argument

type CreateCustomField struct {
	Name      string `json:"name"`
	ChoiceSet string `json:"choice_set"`
	Label     string `json:"label"`
}
