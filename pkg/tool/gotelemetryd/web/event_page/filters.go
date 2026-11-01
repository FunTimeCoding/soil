package event_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Filters(o *query_option.Option) gomponents.Node {
	var chips []gomponents.Node

	for _, name := range []string{
		constant.Tool,
		constant.Surface,
		constant.Actor,
		constant.Kind,
	} {
		value := field(o, name)

		if value == "" {
			continue
		}

		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s: %s", name, value)).
				Link(Path(without(o, name), 1)).
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
