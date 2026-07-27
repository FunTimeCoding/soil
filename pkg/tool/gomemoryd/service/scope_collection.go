package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func ScopeCollection(scope string) string {
	if scope == "" {
		return constant.DefaultCollection
	}

	return scope
}
