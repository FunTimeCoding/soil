package floor_page

import (
	"fmt"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
)

func load(g guest.Guest) string {
	if g.Status != proxmoxConstant.RunningStatus || g.MemoryTotal == 0 {
		return ""
	}

	return fmt.Sprintf(
		"cpu %.0f%% · mem %.0f%%",
		g.Processor*100,
		float64(g.Memory)/float64(g.MemoryTotal)*100,
	)
}
