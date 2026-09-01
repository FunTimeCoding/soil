package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
)

func nodeLabel(n floor.Node) string {
	if n.Hypervisor == n.Name {
		return n.Name
	}

	return fmt.Sprintf("%s · %s", n.Hypervisor, n.Name)
}
