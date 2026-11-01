package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func listScope(scope string) string {
	if scope == "" {
		return constant.AllScope
	}

	if scope == constant.DefaultScope {
		return ""
	}

	return scope
}
