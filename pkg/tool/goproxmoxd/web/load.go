package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
)

func load(g floor.Guest) string {
	if g.Status != constant.RunningStatus || g.MemoryTotal == 0 {
		return ""
	}

	return fmt.Sprintf(
		"cpu %.0f%% · mem %.0f%%",
		g.Processor*100,
		float64(g.Memory)/float64(g.MemoryTotal)*100,
	)
}
