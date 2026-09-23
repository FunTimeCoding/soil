package choice

func New(label string, position int) *Choice {
	return &Choice{Label: label, Position: position}
}
