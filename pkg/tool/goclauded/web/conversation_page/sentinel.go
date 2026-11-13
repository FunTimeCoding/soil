package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sentinel(offset int) gomponents.Node {
	return html.Small(
		html.Class(constant.SidebarMoreClass),
		extended.Get(sidebarPath(offset)),
		extended.Trigger(webConstant.TriggerShown),
		extended.Swap(webConstant.SwapOuter),
		gomponents.Text("Loading more…"),
	)
}
