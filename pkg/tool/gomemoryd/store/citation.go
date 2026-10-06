package store

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func citation(
	scope string,
	name string,
) string {
	if scope == "" {
		scope = constant.DefaultScope
	}

	return join.Empty(
		constant.Backtick,
		constant.MemoryScheme,
		scope,
		stringsConstant.Slash,
		name,
		constant.Backtick,
	)
}
