package entry_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(e *entry.Entry) gomponents.Node {
	r := opening(e)

	return r.Detail(
		constant.EntryColumns,
		fields(e),
		html.Div(
			html.Class(webConstant.ActionRowClass),
			load.New(fragment(constant.EditPath, e.Identifier), "Edit").
				Swap(r.Target(), webConstant.SwapOuter).
				Button().
				Class("outline").
				Node(),
			action.New(fragment(constant.DeletePath, e.Identifier), "Delete").
				Swap(r.Target(), webConstant.SwapOuter).
				Confirm(fmt.Sprintf("Delete entry %d?", e.Identifier)).
				ButtonClass("outline contrast").
				Node(),
			load.New(fragment(constant.ClosePath, e.Identifier), "Close").
				Swap(r.Target(), webConstant.SwapOuter).
				Button().
				Class("outline secondary").
				Node(),
			load.New(Locator(e.Identifier), "Permalink").
				Button().
				Class("outline secondary").
				Node(),
		),
	).Response()
}
