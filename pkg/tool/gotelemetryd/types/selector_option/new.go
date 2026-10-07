package selector_option

func New(
	value string,
	label string,
) *Option {
	return &Option{Value: value, Label: label}
}
