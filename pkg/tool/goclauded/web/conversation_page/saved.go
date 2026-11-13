package conversation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Saved(
	identifier string,
	alias string,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.P(gomponents.Text(fmt.Sprintf("Saved: %s", alias))),
			load.New(panelPath(identifier), "View conversation").
				Swap(constant.PanelTarget, webConstant.SwapInner).
				Node(),
		},
	)
}
