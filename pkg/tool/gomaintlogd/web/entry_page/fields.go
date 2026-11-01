package entry_page

import (
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func fields(e *entry.Entry) gomponents.Node {
	return fact.New().
		Row("Timestamp", time.FormatCompact(e.Timestamp)).
		Row("Action", e.Action).
		Row("User", e.User).
		Row("System", e.System).
		Row("Service", e.Service).
		RowNode("Description", html.Pre(gomponents.Text(e.Description))).
		Node()
}
