package web

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func relatedScope(scope string) string {
	if scope == "" {
		return constant.DefaultScope
	}

	return scope
}
