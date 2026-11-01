package impression_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func windows(current int) gomponents.Node {
	chips := make([]gomponents.Node, 0, len(constant.ImpressionWindows))

	for _, d := range constant.ImpressionWindows {
		c := chip.New(fmt.Sprintf("%dd", d)).Link(Path(d))

		if d == current {
			c.Current()
		}

		chips = append(chips, c.Node())
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Group(chips),
	)
}
