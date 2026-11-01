package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(l *memory_list.List) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.MemoriesTitle)),
			scopes(l),
			filters(l),
			listTable(l),
			listPager(l),
		},
	)
}
