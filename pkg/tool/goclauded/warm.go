package goclauded

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service"

func warm(v *service.Service) {
	for _, e := range v.Sessions() {
		v.ToolCalls(e.Identifier)
	}
}
