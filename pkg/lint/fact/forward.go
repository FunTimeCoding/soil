package fact

type Forward struct {
	Argument   int         `json:"argument"`
	Conditions []Condition `json:"conditions"`
}
