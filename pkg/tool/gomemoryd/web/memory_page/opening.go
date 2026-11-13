package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"maragu.dev/gomponents"
)

func opening(
	identifier int64,
	cells []gomponents.Node,
) *opening_row.OpeningRow {
	return opening_row.New(
		fmt.Sprintf("%s%d", constant.RowPrefix, identifier),
		actionPath(identifier, constant.OpenAction),
		actionPath(identifier, constant.CloseAction),
	).Cells(cells...)
}
