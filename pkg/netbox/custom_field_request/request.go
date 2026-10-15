package custom_field_request

type Request struct {
	Name        string   `json:"name"`
	Label       string   `json:"label,omitempty"`
	Kind        string   `json:"type"`
	ObjectTypes []string `json:"object_types"`
	ChoiceSet   int32    `json:"choice_set"`
}
