package record

import "github.com/funtimecoding/soil/pkg/telemetry/constant"

func validSurface(surface string) bool {
	switch surface {
	case constant.ModelContext, constant.CommandLine, constant.Web, constant.WebService:
		return true
	default:
		return false
	}
}
