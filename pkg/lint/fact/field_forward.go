package fact

type FieldForward struct {
	Argument   int         `json:"argument"`
	Field      string      `json:"field"`
	Conditions []Condition `json:"conditions"`
}
