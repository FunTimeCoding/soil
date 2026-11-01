package field_row

import "maragu.dev/gomponents"

func (r *FieldRow) Actions(n ...gomponents.Node) *FieldRow {
	r.actions = append(r.actions, n...)

	return r
}
