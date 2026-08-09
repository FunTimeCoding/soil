package gomemory

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func scopeLabel(scope *string) string {
	if scope == nil || *scope == "" {
		return constant.DefaultScope
	}

	return *scope
}
