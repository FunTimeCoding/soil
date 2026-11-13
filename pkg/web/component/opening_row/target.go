package opening_row

import "github.com/funtimecoding/soil/pkg/strings/join"

func (r *OpeningRow) Target() string {
	return join.Empty("#", r.detailIdentifier())
}
