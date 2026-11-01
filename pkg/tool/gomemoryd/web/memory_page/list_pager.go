package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"maragu.dev/gomponents"
)

func listPager(l *memory_list.List) gomponents.Node {
	if l.Total == 0 {
		return gomponents.Group(nil)
	}

	offset := (l.Page - 1) * l.Limit
	result := pager.New().Position(
		fmt.Sprintf(
			"%d–%d of %d",
			min(offset+1, l.Total),
			min(offset+l.Limit, l.Total),
			l.Total,
		),
	)

	if l.Page > 1 {
		result.Previous(ListPath(l.Type, l.Tag, l.Scope, l.Page-1), "← Newer")
	}

	if offset+l.Limit < l.Total {
		result.Next(ListPath(l.Type, l.Tag, l.Scope, l.Page+1), "Older →")
	}

	return result.Node()
}
