package record

import "github.com/funtimecoding/soil/pkg/telemetry/constant"

func validKind(kind string) bool {
	switch kind {
	case constant.Baseline, constant.Domain:
		return true
	default:
		return false
	}
}
