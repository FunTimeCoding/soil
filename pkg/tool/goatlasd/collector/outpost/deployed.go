package outpost

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

func deployed(v client.Service) bool {
	if v.Origin == systemConstant.ServiceOriginSystem {
		return false
	}

	return v.Source != nil && *v.Source != ""
}
