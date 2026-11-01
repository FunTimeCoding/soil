package floor_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/node"
)

func label(n node.Node) string {
	if n.Hypervisor == n.Name {
		return n.Name
	}

	return fmt.Sprintf("%s · %s", n.Hypervisor, n.Name)
}
