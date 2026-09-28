package web

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"

func AnyScope(v []*placement.Placement) bool {
	for _, p := range v {
		if p.Scope != "" {
			return true
		}
	}

	return false
}
