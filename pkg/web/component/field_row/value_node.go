package field_row

import "maragu.dev/gomponents"

func (r *FieldRow) ValueNode(n ...gomponents.Node) *FieldRow {
	r.valueNode = append(r.valueNode, n...)

	return r
}
