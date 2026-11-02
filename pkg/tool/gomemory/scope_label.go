package gomemory

import gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func scopeLabel(scope *string) string {
	if scope == nil || *scope == "" {
		return gomemorydConstant.DefaultScope
	}

	return *scope
}
