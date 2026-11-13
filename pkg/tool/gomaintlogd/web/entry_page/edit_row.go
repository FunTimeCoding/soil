package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	component "github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func EditRow(
	e *entry.Entry,
	alert string,
) gomponents.Node {
	r := opening(e)

	return r.Detail(
		constant.EntryColumns,
		component.New(fragment(constant.EditPath, e.Identifier), "Save").
			Swap(r.Target(), webConstant.SwapOuter).
			Alert(alert).
			Text(constant.Action, "Action", e.Action).
			Required().
			Text(constant.User, "User", e.User).
			Required().
			DateTime(
				constant.Timestamp,
				"Timestamp",
				e.Timestamp.Format(constant.FilterFormat),
			).
			Text(constant.System, "System", e.System).
			Text(constant.Service, "Service", e.Service).
			TextArea(constant.Description, "Description", e.Description).
			Lines("4").
			Beside(
				load.New(fragment(constant.DetailPath, e.Identifier), "Cancel").
					Swap(r.Target(), webConstant.SwapOuter).
					Button().
					Class("outline secondary").
					Node(),
			).
			Node(),
	).Response()
}
