package custom_field_record

type Record struct {
	Identifier  int32    `json:"id"`
	Name        string   `json:"name"`
	Label       string   `json:"label"`
	Kind        Kind     `json:"type"`
	ObjectTypes []string `json:"object_types"`
	ChoiceSet   *Named   `json:"choice_set"`
}
