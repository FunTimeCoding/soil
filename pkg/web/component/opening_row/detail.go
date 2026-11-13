package opening_row

import "maragu.dev/gomponents"

func (r *OpeningRow) Detail(
	columns int,
	content ...gomponents.Node,
) *OpeningRow {
	r.columns = columns
	r.detail = content
	r.expanded = true

	return r
}
