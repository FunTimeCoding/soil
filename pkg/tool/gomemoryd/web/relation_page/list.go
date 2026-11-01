package relation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(
	relations []record.RelationOverview,
	filter string,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.RelationsTitle)),
			filters(relations, filter),
			listTable(relations, filter),
		},
	)
}
