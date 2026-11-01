package log_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"maragu.dev/gomponents"
)

func pagination(
	offset int,
	total int,
	start string,
	end string,
) gomponents.Node {
	p := pager.New().Position(
		fmt.Sprintf(
			"%d–%d of %d",
			offset+1,
			min(offset+constant.PageSize, total),
			total,
		),
	)

	if offset > 0 {
		p.Previous(
			pageLink(max(offset-constant.PageSize, 0), start, end),
			"Previous",
		)
	}

	if offset+constant.PageSize < total {
		p.Next(pageLink(offset+constant.PageSize, start, end), "Next")
	}

	return p.Node()
}
