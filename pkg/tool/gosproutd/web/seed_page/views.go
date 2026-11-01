package seed_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Views(modified bool) gomponents.Node {
	priority := chip.New("priority").Link(Path(false))
	recent := chip.New(constant.SortModified).Link(Path(true))

	if modified {
		recent.Current()
	} else {
		priority.Current()
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		priority.Node(),
		recent.Node(),
	)
}
