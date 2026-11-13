package opening_row

import "maragu.dev/gomponents"

func (r *OpeningRow) Node() gomponents.Node {
	return gomponents.Group([]gomponents.Node{r.row(), r.detailRow()})
}
