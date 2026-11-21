package system_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func NewSummary(v *system.System) *Summary {
	return &Summary{
		Name:         v.Name,
		Release:      v.Release,
		Kernel:       v.Kernel,
		Priority:     v.Priority,
		Labels:       v.Labels,
		Coverage:     constant.CoverageUntracked,
		Updated:      v.UpdatedAt,
		ListsUpdated: v.ListsUpdated,
		KernelOffers: v.KernelOffers,
	}
}
