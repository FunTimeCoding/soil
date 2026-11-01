package impression_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(
	impressions []record.Impression,
	days int,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.ImpressionsTitle)),
			windows(days),
			listTable(impressions),
		},
	)
}
