package series

func New(
	label string,
	class string,
) *Series {
	return &Series{Label: label, Class: class}
}
