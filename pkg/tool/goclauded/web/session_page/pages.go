package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"maragu.dev/gomponents"
)

func pages(
	active bool,
	page int,
	total int,
) gomponents.Node {
	if total == 0 {
		return gomponents.Group(nil)
	}

	offset := (page - 1) * constant.PageSize
	result := pager.New().Position(
		fmt.Sprintf(
			"%d–%d of %d",
			min(offset+1, total),
			min(offset+constant.PageSize, total),
			total,
		),
	)

	if page > 1 {
		result.Previous(ListPath(active, page-1), "← Newer")
	}

	if offset+constant.PageSize < total {
		result.Next(ListPath(active, page+1), "Older →")
	}

	return result.Node()
}
