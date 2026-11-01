package entry_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/url"
)

func Filter(values url.Values) gomponents.Node {
	return entry.New(constant.EntriesPath, "Filter").
		Query().
		Swap(join.Empty("#", constant.EntriesMark), webConstant.SwapInner).
		Inline().
		Text(constant.System, "System", values.Get(constant.System)).
		Placeholder("e.g. worker1").
		Text(constant.Service, "Service", values.Get(constant.Service)).
		Placeholder("e.g. nginx").
		Text(constant.User, "User", values.Get(constant.User)).
		DateTime(constant.Since, "Since", values.Get(constant.Since)).
		DateTime(constant.Until, "Until", values.Get(constant.Until)).
		Node()
}
