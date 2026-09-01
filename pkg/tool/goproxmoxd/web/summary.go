package web

import (
	"fmt"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
)

func summary(f floor.Floor) []string {
	running := 0
	unbacked := 0

	for _, g := range f.Guests {
		if g.Status == proxmoxConstant.RunningStatus {
			running++
		}

		if g.Unbacked {
			unbacked++
		}
	}

	updates := 0

	for _, n := range f.Nodes {
		updates += n.UpdatesPending
	}

	result := []string{
		fmt.Sprintf("%d guests", len(f.Guests)),
		fmt.Sprintf("%d running", running),
	}

	if updates > 0 {
		result = append(result, fmt.Sprintf("%d updates", updates))
	}

	if unbacked > 0 {
		result = append(result, fmt.Sprintf("%d unbacked", unbacked))
	}

	for _, t := range f.Storages {
		if t.Total == 0 {
			continue
		}

		ratio := float64(t.Used) / float64(t.Total)

		if ratio >= constant.StorageWarnRatio {
			result = append(result, fmt.Sprintf("%s %.0f%%", t.Name, ratio*100))
		}
	}

	return result
}
