package relation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func filters(
	relations []record.RelationOverview,
	current string,
) gomponents.Node {
	counts := map[string]int{}

	for _, v := range relations {
		if v.Type == "" {
			counts[constant.UntypedFilter]++
		} else {
			counts[v.Type]++
		}
	}

	chips := []gomponents.Node{
		filterChip(fmt.Sprintf("all (%d)", len(relations)), "", current),
		filterChip(
			fmt.Sprintf("untyped (%d)", counts[constant.UntypedFilter]),
			constant.UntypedFilter,
			current,
		),
	}

	for _, name := range constant.RelationTypes {
		if counts[name] > 0 {
			chips = append(
				chips,
				filterChip(
					fmt.Sprintf("%s (%d)", name, counts[name]),
					name,
					current,
				),
			)
		}
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Text("Type:"),
		gomponents.Group(chips),
	)
}
