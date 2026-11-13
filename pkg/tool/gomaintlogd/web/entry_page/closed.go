package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
)

func Closed(e *entry.Entry) gomponents.Node {
	return opening(e).Response()
}
