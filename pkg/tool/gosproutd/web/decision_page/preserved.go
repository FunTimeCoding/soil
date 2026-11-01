package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
)

func Preserved(d *decision.Decision) gomponents.Node {
	return compact(d).Preserve().Node()
}
