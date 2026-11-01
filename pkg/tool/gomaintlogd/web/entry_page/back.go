package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	"maragu.dev/gomponents"
)

func back() gomponents.Node {
	return load.New(constant.EntriesPath, constant.EntriesTitle).
		Button().
		Class("outline secondary").
		Node()
}
