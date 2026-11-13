package opening_row

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (r *OpeningRow) detailIdentifier() string {
	return join.Empty(r.identifier, constant.OpeningDetailSuffix)
}
