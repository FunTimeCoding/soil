package fact

type Release struct {
	Field     string    `json:"field"`
	Condition Condition `json:"condition"`
}
