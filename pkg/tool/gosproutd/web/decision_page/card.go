package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
)

func Card(d *decision.Decision) gomponents.Node {
	return compact(d).Node()
}
