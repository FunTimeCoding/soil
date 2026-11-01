package field_row

import "maragu.dev/gomponents"

type FieldRow struct {
	label      string
	identifier string
	tooltip    string
	value      string
	note       string
	valueNode  []gomponents.Node
	actions    []gomponents.Node
}
