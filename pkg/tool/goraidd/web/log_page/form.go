package log_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
)

func Form(listing gomponents.Node) gomponents.Node {
	return entry.New(constant.GeneratePath, "Generate Report").
		Content(listing).
		Alternative(constant.CreateRaidPath, "Create Raid").
		Node()
}
