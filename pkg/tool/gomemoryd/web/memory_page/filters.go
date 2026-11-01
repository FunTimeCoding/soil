package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func filters(l *memory_list.List) gomponents.Node {
	var chips []gomponents.Node

	if l.Tag != "" {
		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s: %s", constant.Tag, l.Tag)).
				Link(ListPath(l.Type, "", l.Scope, 1)).
				Content(gomponents.Text(" ×")).
				Node(),
		)
	}

	if l.Type != "" {
		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s: %s", constant.Type, l.Type)).
				Link(ListPath("", l.Tag, l.Scope, 1)).
				Content(gomponents.Text(" ×")).
				Node(),
		)
	}

	if len(chips) == 0 {
		return gomponents.Group(nil)
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Text("Filtered:"),
		gomponents.Group(chips),
	)
}
