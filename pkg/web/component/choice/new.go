package choice

func New(
	value string,
	label string,
) *Choice {
	return &Choice{Value: value, Label: label}
}
