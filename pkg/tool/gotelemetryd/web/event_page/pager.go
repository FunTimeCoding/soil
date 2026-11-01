package event_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"maragu.dev/gomponents"
)

func Pager(
	o *query_option.Option,
	page int,
	more bool,
) gomponents.Node {
	p := pager.New()

	if page > 1 {
		p.Previous(Path(o, page-1), "← Newer")
	}

	if more {
		p.Next(Path(o, page+1), "Older →")
	}

	return p.Node()
}
