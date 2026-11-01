package log_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
)

func Filter(
	start string,
	end string,
) gomponents.Node {
	return entry.New(constant.LogsPath, "Filter").
		Query().
		Inline().
		DateTime(constant.StartParameter, "From", start).
		DateTime(constant.EndParameter, "To", end).
		Node()
}
