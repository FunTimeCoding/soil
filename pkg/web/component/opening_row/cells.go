package opening_row

import "maragu.dev/gomponents"

func (r *OpeningRow) Cells(cells ...gomponents.Node) *OpeningRow {
	r.cells = append(r.cells, cells...)

	return r
}
