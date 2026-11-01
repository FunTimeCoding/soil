package component

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func ScopeName(scope string) string {
	if scope == "" {
		return constant.DefaultScope
	}

	return scope
}
