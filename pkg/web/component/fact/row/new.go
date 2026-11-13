package row

import "maragu.dev/gomponents"

func New(
	label string,
	value gomponents.Node,
) *Row {
	return &Row{Label: label, Value: value}
}
