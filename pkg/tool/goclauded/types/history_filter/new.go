package history_filter

func New(
	label string,
	kind string,
) *Filter {
	return &Filter{Label: label, Kind: kind}
}
