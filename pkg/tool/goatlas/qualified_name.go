package goatlas

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

func qualifiedName(v client.Placement) string {
	if v.Scope == nil || *v.Scope == "" {
		return v.Name
	}

	return join.Empty(*v.Scope, constant.ScopeSeparator, v.Name)
}
