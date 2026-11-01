package entry_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/url"
)

func AddForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.AddEntryPath, constant.AddEntryTitle).
		Identifier(constant.AddFormMark).
		Swap(join.Empty("#", constant.AddFormMark), webConstant.SwapOuter).
		Alert(alert).
		Text(constant.Action, "Action (required)", values.Get(constant.Action)).
		Required().
		Placeholder("e.g. restarted web server").
		Text(constant.User, "User (required)", values.Get(constant.User)).
		Required().
		Placeholder("e.g. jdoe").
		Text(constant.System, "System", values.Get(constant.System)).
		Placeholder("e.g. worker1").
		Text(constant.Service, "Service", values.Get(constant.Service)).
		Placeholder("e.g. nginx").
		TextArea(
			constant.Description,
			"Description",
			values.Get(constant.Description),
		).
		Placeholder("Optional details...").
		Lines("3").
		Node()
}
