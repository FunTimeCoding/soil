package host_mapping

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"slices"
)

func action(actions []string) string {
	create := slices.Contains(actions, constant.PlanCreate)
	remove := slices.Contains(actions, constant.PlanDelete)

	switch {
	case create && remove:
		return provisionConstant.HostReplace
	case create:
		return provisionConstant.HostCreate
	case remove:
		return provisionConstant.HostDelete
	case slices.Contains(actions, constant.PlanUpdate):
		return provisionConstant.HostUpdate
	default:
		return ""
	}
}
